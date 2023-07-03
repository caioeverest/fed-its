package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/caioeverest/fed-its/adapter/database"
	"github.com/caioeverest/fed-its/adapter/redis"
	"github.com/caioeverest/fed-its/internal/aes"
	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/itserrors"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/caioeverest/fed-its/internal/validate"
	"github.com/caioeverest/fed-its/model"
	"github.com/samber/lo"
)

const hide = "***********"

type ProviderI interface {
	Create(ctx context.Context, provider model.Provider) (model.Provider, error)
	Get(ctx context.Context, slug string) (model.Provider, error)
	Update(ctx context.Context, signature, slug string, provider model.Provider) (model.Provider, error)
	Delete(ctx context.Context, signature, slug string) error
	List(ctx context.Context, method string) ([]model.Provider, error)
}

type Proveder struct {
	cfg      *config.Config
	log      *logger.Logger
	db       *database.Database
	validate *validate.Validate
	redis    *redis.Client
}

func NewProvider(cfg *config.Config, log *logger.Logger, db *database.Database, validate *validate.Validate, redis *redis.Client) ProviderI {
	return &Proveder{cfg, log, db, validate, redis}
}

// Create godoc
// @Summary Create a new provider
// @Description Create a new provider
func (p *Proveder) Create(ctx context.Context, provider model.Provider) (result model.Provider, err error) {
	p.log.Info("New provider requested to be created")

	//Validate input
	p.log.Info("Validating provider")
	if err = p.validate.Struct(provider); err != nil {
		p.log.Errorf("Validation error: %+v", err)
		return
	}

	//Encrypt secret
	p.log.Info("Encrypting provider secret")
	if err = p.encrypt(ctx, &provider); err != nil {
		p.log.Errorf("Error encrypting provider secret - %+v", err)
		return
	}

	//Create provider
	p.log.Infof("Creating %s provider", provider.Slug)
	if err = p.db.WithContext(ctx).Create(&provider).Error; err != nil {
		p.log.Errorf("Error creating provider - %+v", err)
		return
	}
	p.log.Info("Provider created")
	return provider, nil
}

// Get godoc
// @Summary Get a provider
// @Description Get a provider by slug name and return it
func (p *Proveder) Get(ctx context.Context, slug string) (provider model.Provider, err error) {
	p.log.Infof("Get provider with slug %s", slug)
	if err = p.db.WithContext(ctx).Where("slug = ?", slug).First(&provider).Error; err != nil {
		p.log.Errorf("Error getting provider - %+v", err)
		return
	}
	provider.Secret = hide

	p.log.Infof("Provider %s found!", slug)
	return
}

// Update godoc
// @Summary Update a provider
// @Description Update a provider by slug name and return it
func (p *Proveder) Update(ctx context.Context, signature, slug string, update model.Provider) (provider model.Provider, err error) {
	p.log.Infof("Update provider %s requested", slug)

	//Search for provider
	p.log.Infof("Searching for provider %s", slug)
	if err = p.db.WithContext(ctx).Where("slug = ?", slug).First(&provider).Error; err != nil {
		p.log.Errorf("Error getting provider - %+v", err)
		return
	}

	//Check signature
	if !p.checkSignature(ctx, provider, signature, provider) {
		p.log.Errorf("Signature check failed")
		return model.Provider{}, itserrors.ErrInvalidSignature
	}

	//Update provider
	if err = p.db.WithContext(ctx).Model(&provider).Updates(update).Error; err != nil {
		p.log.Errorf("Error updating provider - %+v", err)
		return
	}
	provider.Secret = hide

	return provider, nil
}

// Delete godoc
// @Summary Delete a provider
// @Description Delete a provider by slug name
func (p *Proveder) Delete(ctx context.Context, signature, slug string) (err error) {
	var (
		provider model.Provider
	)
	p.log.Infof("Delete provider %s requested", slug)

	//Search for provider
	p.log.Infof("Searching for provider %s", slug)
	if err = p.db.WithContext(ctx).Where("slug = ?", slug).First(&provider).Error; err != nil {
		p.log.Errorf("Error getting provider - %+v", err)
		return
	}

	//Check signature
	if !p.checkSignature(ctx, provider, signature, provider) {
		p.log.Errorf("Signature check failed")
		return itserrors.ErrInvalidSignature
	}

	//Delete provider
	if err = p.db.WithContext(ctx).Delete(&provider).Error; err != nil {
		p.log.Errorf("Error deleting provider - %+v", err)
		return
	}

	return nil
}

// List godoc
// @Summary List providers
// @Description List providers that implement a method
func (p *Proveder) List(ctx context.Context, method string) (list []model.Provider, err error) {
	p.log.Infof("List providers that implement method %s", method)

	//List providers
	if err = p.db.WithContext(ctx).
		Joins("JOIN provider_methods ON provider_methods.provider_id = providers.id").
		Joins("JOIN methods ON methods.id = provider_methods.method_id").
		Where("methods.name = ?", method).
		Find(&list).Error; err != nil {
		p.log.Errorf("Error listing providers - %+v", err)
		return
	}
	list = lo.Map(list, func(provider model.Provider, _ int) model.Provider { provider.Secret = hide; return provider })

	p.log.Infof("Found %d providers that implement %s", len(list), method)
	return
}

func (p *Proveder) checkSignature(ctx context.Context, model model.Provider, signature string, content any) bool {
	var (
		contentBytes []byte
		secret       string
		err          error
	)
	p.log.Infof("Checking signature for provider %s", model.Slug)
	if secret, err = p.decrypt(ctx, model); err != nil {
		p.log.Errorf("Error decrypting provider secret - %+v", err)
		return false
	}
	if contentBytes, err = json.Marshal(content); err != nil {
		p.log.Errorf("Error marshaling content - %+v", err)
		return false
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(contentBytes))
	sha := hex.EncodeToString(h.Sum(nil))

	return sha == signature
}

func (p *Proveder) encrypt(ctx context.Context, model *model.Provider) (err error) {
	model.Secret, err = aes.Encrypt(p.cfg.HashSecret, model.Secret)
	return
}

func (p *Proveder) decrypt(ctx context.Context, model model.Provider) (secret string, err error) {
	secret, err = aes.Decrypt(p.cfg.HashSecret, model.Secret)
	return
}
