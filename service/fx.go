package service

import "go.uber.org/fx"

func Services() fx.Option {
	return fx.Provide(
		NewProvider,
		NewMethod,
	)
}
