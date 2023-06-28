package handler

import (
	"github.com/caioeverest/fedits/internal/config"
	"github.com/caioeverest/fedits/internal/logger"
	"github.com/caioeverest/fedits/model"
	"github.com/caioeverest/fedits/service"
	"github.com/labstack/echo/v4"
)

type Method struct {
	cfg     *config.Config
	log     *logger.Logger
	service service.MethodI
}

func NewMethod(cfg *config.Config, log *logger.Logger, service service.MethodI) *Method {
	handler := &Method{
		cfg:     cfg,
		log:     log,
		service: service,
	}
	return handler
}

// Create godoc
// @Summary Create a new method
// @Description Create a new method
// @Tags method
// @Accept json
// @Produce json
// @Param payload body model.Method true "payload"
// @Success 200 {object} model.Method
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /method [post]
func (m *Method) Create(pctx echo.Context) (err error) {
	var (
		payload model.Method
		result  model.Method
		ctx     = pctx.Request().Context()
	)

	if err = pctx.Bind(&payload); err != nil {
		m.log.Errorf("Error binding payload: %v", err)
		return
	}

	if result, err = m.service.Create(ctx, payload); err != nil {
		m.log.Errorf("Error creating method: %v", err)
		return
	}

	return pctx.JSON(201, result)
}

// List godoc
// @Summary List methods
// @Description List methods
// @Tags method
// @Accept json
// @Produce json
// @Success 200 {array} model.Method
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /method [get]
func (m *Method) List(pctx echo.Context) (err error) {
	var (
		result []model.Method
		ctx    = pctx.Request().Context()
	)

	if result, err = m.service.List(ctx); err != nil {
		m.log.Errorf("Error listing methods: %v", err)
		return
	}

	return pctx.JSON(200, result)
}

// Get godoc
// @Summary Get method
// @Description Get method
// @Tags method
// @Accept json
// @Produce json
// @Param method path string true "Method"
// @Success 200 {object} model.Method
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /method/{method} [get]
func (m *Method) Get(pctx echo.Context) (err error) {
	var (
		result model.Method
		method = pctx.Param("method")
		ctx    = pctx.Request().Context()
	)

	if result, err = m.service.Get(ctx, method); err != nil {
		m.log.Errorf("Error getting method %s: %v", method, err)
		return
	}

	return pctx.JSON(200, result)
}
