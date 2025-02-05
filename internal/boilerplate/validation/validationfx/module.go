package validationfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/validation"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"validation",
		fx.Provide(validation.New),
	)
}
