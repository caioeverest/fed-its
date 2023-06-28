package service

import (
	"context"

	"github.com/caioeverest/fedits/adapter/database"
	"github.com/caioeverest/fedits/internal/config"
	"github.com/caioeverest/fedits/internal/logger"
	"github.com/caioeverest/fedits/model"
)

type Orquestrate interface {
	Request(ctx context.Context, userRef string, method string, params []any) (result model.Envelope, err error)
}

type Orquestrator struct {
	conf *config.Config
	log  *logger.Logger
	db   *database.Database
}

func NewOrquestrator(conf *config.Config, log *logger.Logger, db *database.Database) Orquestrate {
	return &Orquestrator{conf, log, db}
}

// Request godoc
// @Summary Request a method
// @Description Request a method from a provider or a group of providers and return the first response received
func (o *Orquestrator) Request(ctx context.Context, userRef string, methodName string, params []any) (result model.Envelope, err error) {
	var (
		method          model.Method
		listOfProviders []model.Provider
	)
	o.log.Infof("New request received from user %s to call method %s", userRef, methodName)
	if err = o.db.WithContext(ctx).Where("name = ?", methodName).First(&method).Error; err != nil {
		o.log.Errorf("Error while validating request: %+v", err)
		return
	}

	// Validate input
	o.log.Info("Validating request")
	// TODO: Validate input

	// Get list of providers
	o.log.Info("Getting list of providers")
	if err = o.db.WithContext(ctx).
		Joins("JOIN provider_methods ON provider_methods.provider_id = providers.id").
		Joins("JOIN methods ON methods.id = provider_methods.method_id").
		Where("methods.name = ?", method).
		Find(&listOfProviders).Error; err != nil {
		o.log.Errorf("Error while validating request: %+v", err)
		return
	}
	o.log.Infof("Found %d providers", len(listOfProviders))

	// Create a new context, with cancellation capabilities
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Use channels to collect the results
	var (
		resultsChan = make(chan model.Envelope, len(listOfProviders))
		errorsChan  = make(chan error, len(listOfProviders))
	)

	// Call providers
	for _, provider := range listOfProviders {
		go o.callProvider(ctx, cancel, provider, resultsChan, errorsChan, userRef, methodName, params)
	}

	// Wait for the first response
	select {
	case result = <-resultsChan:
		o.log.Infof("Got a response from provider")
	case err = <-errorsChan:
		o.log.Errorf("Got an error from provider: %+v", err)
	}

	return
}

// callProvider launch a goroutine for each provider
func (o *Orquestrator) callProvider(ctx context.Context, closeRun func(), provider model.Provider, resultsChan chan model.Envelope, errorsChan chan error, userRef, methodName string, params []any) {
	// Here you'll need to replace 'CallProviderMethod' with the actual method that makes the request
	if response, err := provider.CallProviderMethod(ctx, o.conf.HashSecret, userRef, methodName, params); err != nil {
		errorsChan <- err
	} else {
		resultsChan <- model.Envelope{
			Provider: provider.Name,
			Result:   response.SuccessResult(),
		} // Push the response into the results channel
		closeRun() // Cancel the context, this will stop other running requests
	}
}
