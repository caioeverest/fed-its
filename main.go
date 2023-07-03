package main

import (
	"context"

	"github.com/caioeverest/fed-its/adapter/database"
	"github.com/caioeverest/fed-its/adapter/http"
	"github.com/caioeverest/fed-its/adapter/redis"
	"github.com/caioeverest/fed-its/handler"
	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/caioeverest/fed-its/internal/validate"
	"github.com/caioeverest/fed-its/model"
	"github.com/caioeverest/fed-its/service"
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
// @BasePath       /
func main() {
	app := fx.New(
		handler.Providers(),
		handler.Invoke(),
		service.Services(),
		fx.Provide(http.New, database.New, redis.New),
		fx.Provide(validate.New, logger.New, config.New),
		fx.Invoke(model.Migrate, func(*http.Server) {}),
	)
	defer close(app)
	app.Run()
}

func close(app *fx.App) {
	if err := app.Stop(context.Background()); err != nil {
		panic(err)
	}
}
