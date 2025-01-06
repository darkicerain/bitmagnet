package importerfx

import (
	"bitmagnet-io/bitmagnet/internal/importer"
	"bitmagnet-io/bitmagnet/internal/importer/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"importer",
		fx.Provide(
			httpserver.New,
			importer.New,
		),
	)
}
