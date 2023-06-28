package handler

import (
	"github.com/caioeverest/fedits/internal/config"
	_ "github.com/caioeverest/fedits/internal/itserrors"
	"github.com/caioeverest/fedits/internal/logger"
	"github.com/caioeverest/fedits/model"
	"github.com/caioeverest/fedits/service"
	"github.com/labstack/echo/v4"
)

type Provider struct {
	cfg     *config.Config
	log     *logger.Logger
	service service.ProviderI
}

func NewProvider(cfg *config.Config, log *logger.Logger, service service.ProviderI) *Provider {
	handler := &Provider{
		cfg:     cfg,
		log:     log,
		service: service,
	}

	return handler
}

// Create godoc
// @Summary Create a new provider
// @Description Create a new provider
// @Tags provider
// @Accept json
// @Produce json
// @Param provider body model.Provider true "Provider"
// @Success 200 {object} model.Provider
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /provider [post]
func (p *Provider) Create(pctx echo.Context) (err error) {
	var (
		payload model.Provider
		result  model.Provider
		ctx     = pctx.Request().Context()
	)

	if err = pctx.Bind(&payload); err != nil {
		p.log.Errorf("Error binding payload: %v", err)
		return
	}

	if result, err = p.service.Create(ctx, payload); err != nil {
		p.log.Errorf("Error creating provider: %v", err)
		return
	}

	return pctx.JSON(201, result)
}

// Get godoc
// @Summary Get a provider
// @Description Get a provider
// @Tags provider
// @Accept json
// @Produce json
// @Param slug path string true "Provider slug"
// @Success 200 {object} model.Provider
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /provider/{slug} [get]
func (p *Provider) Get(pctx echo.Context) (err error) {
	var (
		result model.Provider
		ctx    = pctx.Request().Context()
		slug   = pctx.Param("slug")
	)

	if result, err = p.service.Get(ctx, slug); err != nil {
		p.log.Errorf("Error getting provider: %v", err)
		return
	}

	return pctx.JSON(200, result)
}

// Update godoc
// @Summary Update a provider
// @Description Update a provider
// @Tags provider
// @Accept json
// @Produce json
// @Param slug path string true "Provider slug"
// @Param provider body model.Provider true "Provider"
// @Param X-Signature header string true "Signature"
// @Success 200 {object} Provider
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /provider/{slug} [patch]
func (p *Provider) Update(pctx echo.Context) (err error) {
	var (
		payload   model.Provider
		result    model.Provider
		slug      = pctx.Param("slug")
		ctx       = pctx.Request().Context()
		signature = pctx.Request().Header.Get("X-Signature")
	)

	if err = pctx.Bind(&payload); err != nil {
		p.log.Errorf("Error binding payload: %v", err)
		return
	}
	if result, err = p.service.Update(ctx, signature, slug, payload); err != nil {
		p.log.Errorf("Error update provider: %v", err)
		return
	}

	return pctx.JSON(200, result)
}

// Delete godoc
// @Summary Delete a provider
// @Description Delete a provider
// @Tags provider
// @Accept json
// @Produce json
// @Param slug path string true "Provider slug"
// @Success 200 {object} model.Provider
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /provider/{slug} [delete]
func (p *Provider) Delete(pctx echo.Context) (err error) {
	var (
		ctx       = pctx.Request().Context()
		slug      = pctx.Param("slug")
		signature = pctx.Request().Header.Get("X-Signature")
	)

	if err = p.service.Delete(ctx, signature, slug); err != nil {
		p.log.Errorf("Error delete provider: %v", err)
		return
	}

	return pctx.JSON(200, nil)
}

// List godoc
// @Summary List providers
// @Description List providers
// @Tags provider
// @Accept json
// @Produce json
// @Param method path string true "Provider method"
// @Success 200 {object} model.Provider
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /provider/list/{method} [get]
func (p *Provider) List(pctx echo.Context) (err error) {
	var (
		result []model.Provider
		ctx    = pctx.Request().Context()
		method = pctx.Param("method")
	)

	if result, err = p.service.List(ctx, method); err != nil {
		p.log.Errorf("Error list provider: %v", err)
		return
	}

	return pctx.JSON(200, result)
}
