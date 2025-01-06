package boilerplatefx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/cli/clifx"
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/boilerplate/logging/loggingfx"
	"bitmagnet-io/bitmagnet/internal/boilerplate/validation/validationfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"boilerplate",
		clifx.New(),
		configfx.New(),
		loggingfx.New(),
		validationfx.New(),
	)
}
