package model

import (
	"context"

	"github.com/caioeverest/fedits/adapter/database"
	"go.uber.org/fx"
)

func Migrate(lc fx.Lifecycle, db *database.Database) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				if err = db.AutoMigrate(&Provider{}); err != nil {
					return err
				}
				if err = db.AutoMigrate(&Method{}); err != nil {
					return err
				}
				if err = db.AutoMigrate(&MethodProvider{}); err != nil {
					return err
				}
				return
			},
		},
	)
}
