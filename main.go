package main

import (
	"context"

	"github.com/caioeverest/fedits/adapter/database"
	"github.com/caioeverest/fedits/adapter/http"
	"github.com/caioeverest/fedits/handler"
	"github.com/caioeverest/fedits/internal/config"
	"github.com/caioeverest/fedits/internal/logger"
	"github.com/caioeverest/fedits/internal/validate"
	"github.com/caioeverest/fedits/model"
	"github.com/caioeverest/fedits/service"
	"go.uber.org/fx"
)

// @title          FED ITS API
// @version        1.0
// @description    This is a conceptual API that manages providers, users and methods for the FED ITS PoC.
// @contact.name   Caio Everest
// @contact.email  caioeverest@edu.unirio.br
// @license.name   MIT
// @license.url    https://github.com/caioeverest/fed-its/license
// @host           localhost:8080
// @BasePath  /provider
func main() {
	app := fx.New(
		handler.Providers(),
		handler.Invoke(),
		service.Services(),
		fx.Provide(
			http.New,
			database.New,
			validate.New,
			logger.New,
			config.New,
		),
		fx.Invoke(
			model.Migrate,
			func(*http.Server) {},
		),
	)
	defer func() {
		if err := app.Stop(context.Background()); err != nil {
			panic(err)
		}
	}()
	app.Run()
}
