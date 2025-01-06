package tmdbfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/tmdb"
	"bitmagnet-io/bitmagnet/internal/tmdb/tmdb_health"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"tmdb",
		configfx.NewConfigModule[tmdb.Config]("tmdb", tmdb.NewDefaultConfig()),
		fx.Provide(
			tmdb.New,
			tmdb_health.New,
		),
	)
}
