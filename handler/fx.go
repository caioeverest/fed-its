package handler

import "go.uber.org/fx"

func Providers() fx.Option {
	return fx.Provide(
		NewProvider,
		NewMethod,
		NewOrquestrator,
	)
}

func Invoke() fx.Option {
	return fx.Invoke(
		NewRouter,
	)
}
