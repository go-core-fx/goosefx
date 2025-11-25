package goosefx

import (
	"context"
	"fmt"

	"github.com/go-core-fx/logger"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Module() fx.Option {
	return fx.Module(
		"goosefx",
		logger.WithNamedLogger("goosefx"),
		fx.Provide(NewProvider),
		fx.Invoke(applyMigrations),
	)
}

func applyMigrations(lc fx.Lifecycle, provider *goose.Provider, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("applying migrations")
			res, err := provider.Up(ctx)
			if err != nil {
				return fmt.Errorf("goose up: %w", err)
			}
			log.Info("migrations applied", zap.Int("count", len(res)))
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Info("migrations shutdown completed")
			return nil
		},
	})
}
