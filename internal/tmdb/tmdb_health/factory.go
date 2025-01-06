package tmdb_health

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/lazy"
	"bitmagnet-io/bitmagnet/internal/health"
	"bitmagnet-io/bitmagnet/internal/tmdb"
	"go.uber.org/fx"
	"time"
)

type Params struct {
	fx.In
	Config tmdb.Config
	Client lazy.Lazy[tmdb.Client]
}

type Result struct {
	fx.Out
	Option health.CheckerOption `group:"health_check_options"`
}

func New(p Params) Result {
	return Result{
		Option: health.WithPeriodicCheck(
			time.Minute*5,
			0,
			NewCheck(p.Config.Enabled, p.Client),
		),
	}
}
