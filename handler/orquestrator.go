package handler

import (
	"net/http"

	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/caioeverest/fed-its/model"
	"github.com/caioeverest/fed-its/service"
	"github.com/labstack/echo/v4"
)

type Orquestrator struct {
	conf    *config.Config
	log     *logger.Logger
	service service.Orquestrate
}

func NewOrquestrator(conf *config.Config, log *logger.Logger, service service.Orquestrate) *Orquestrator {
	return &Orquestrator{conf, log, service}
}

type CallRequest struct {
	Method string `json:"method"`
	Params []any  `json:"params"`
}

// Request godoc
// @Summary Request a method
// @Description Request a method from a provider or a group of providers and return the first response received
// @Tags orquestrator
// @Accept json
// @Produce json
// @Param payload body handler.CallRequest true "Payload"
// @Success 200 {object} string
// @Failure      400  {object}  itserrors.Error
// @Failure      404  {object}  itserrors.Error
// @Failure      500  {object}  itserrors.Error
// @Router /call [post]
func (o *Orquestrator) Request(pctx echo.Context) (err error) {
	var (
		ctx    = pctx.Request().Context()
		body   CallRequest
		result model.Envelope
	)

	if err = pctx.Bind(&body); err != nil {
		o.log.Errorf("Error binding payload: %v", err)
		return
	}
	if result, err = o.service.Request(ctx, "", body.Method, body.Params); err != nil {
		o.log.Errorf("Error while validating request: %+v", err)
		return
	}

	return pctx.JSON(http.StatusOK, result)
}
