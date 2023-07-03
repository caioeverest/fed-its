package service

import (
	"context"
	"errors"

	"github.com/caioeverest/fed-its/adapter/database"
	"github.com/caioeverest/fed-its/adapter/redis"
	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/caioeverest/fed-its/model"
	"github.com/imroc/req/v3"
)

type Orquestrate interface {
	Request(ctx context.Context, userRef string, method string, params []any) (result model.Envelope, err error)
}

type Orquestrator struct {
	conf  *config.Config
	log   *logger.Logger
	db    *database.Database
	redis *redis.Client
}

func NewOrquestrator(conf *config.Config, log *logger.Logger, db *database.Database, redis *redis.Client) Orquestrate {
	return &Orquestrator{conf, log, db, redis}
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

	switch method.Kind {
	case model.Broadcast:
		return o.handleBroadcast(ctx, method, listOfProviders, userRef, params)
	case model.Concurrent:
		return o.handleConcurrent(ctx, method, listOfProviders, userRef, params)
	case model.Exchange, model.Indepotent:
		return o.handleIndepotent(ctx, method, listOfProviders, userRef, params)
	default:
		return result, errors.New("method kind not implemented")
	}
}

func (o *Orquestrator) handleBroadcast(ctx context.Context, method model.Method, listOfProviders []model.Provider, userRef string, params []any) (result model.Envelope, err error) {
	var (
		resultsChan = make(chan model.Envelope, len(listOfProviders))
		errorsChan  = make(chan error, len(listOfProviders))
	)

	// Call providers
	for _, provider := range listOfProviders {
		go o.callProvider(ctx, func() {}, provider, resultsChan, errorsChan, userRef, method.Name, params)
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

func (o *Orquestrator) handleConcurrent(pctx context.Context, method model.Method, listOfProviders []model.Provider, userRef string, params []any) (result model.Envelope, err error) {
	var (
		ctx, cancel = context.WithCancel(pctx)
		resultsChan = make(chan model.Envelope, len(listOfProviders))
		errorsChan  = make(chan error, len(listOfProviders))
	)
	defer cancel()

	// Call providers
	for _, provider := range listOfProviders {
		go o.callProvider(ctx, cancel, provider, resultsChan, errorsChan, userRef, method.Name, params)
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

func (o *Orquestrator) handleIndepotent(ctx context.Context, method model.Method, listOfProviders []model.Provider, userRef string, params []any) (result model.Envelope, err error) {
	for _, provider := range listOfProviders {
		var response *req.Response
		if response, err = provider.CallProviderMethod(ctx, o.conf.HashSecret, userRef, method.Name, params); err != nil {
			o.log.Errorf("Got an error from provider: %+v", err)
			continue
		}
		return model.Envelope{
			Provider: provider.Name,
			Result:   response.SuccessResult(),
		}, nil
	}

	return result, errors.New("no provider could handle the request")
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
