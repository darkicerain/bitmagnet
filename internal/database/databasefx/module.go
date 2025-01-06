package databasefx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/database"
	"bitmagnet-io/bitmagnet/internal/database/cache"
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/database/healthcheck"
	"bitmagnet-io/bitmagnet/internal/database/migrations"
	"bitmagnet-io/bitmagnet/internal/database/postgres"
	"bitmagnet-io/bitmagnet/internal/database/search"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"database",
		configfx.NewConfigModule[postgres.Config]("postgres", postgres.NewDefaultConfig()),
		configfx.NewConfigModule[cache.Config]("gorm_cache", cache.NewDefaultConfig()),
		fx.Provide(
			cache.NewInMemoryCacher,
			cache.NewPlugin,
			dao.New,
			database.New,
			healthcheck.New,
			migrations.New,
			postgres.New,
			search.New,
		),
		fx.Decorate(
			cache.NewDecorator,
		),
	)
}
