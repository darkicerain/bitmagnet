package boilerplateappfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/app/cmd/config"
	"bitmagnet-io/bitmagnet/internal/boilerplate/app/cmd/worker"
	"bitmagnet-io/bitmagnet/internal/boilerplate/boilerplatefx"
	"bitmagnet-io/bitmagnet/internal/boilerplate/cli/hooks"
	"bitmagnet-io/bitmagnet/internal/boilerplate/worker/workerfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"app_boilerplate",
		boilerplatefx.New(),
		workerfx.New(),
		fx.Provide(
			hooks.New,
			configcmd.New,
			workercmd.New,
		),
	)
}
