package downloaderfx

import (
	"bitmagnet-io/bitmagnet/internal/downloader/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"downloader",
		fx.Provide(
			httpserver.New,
		),
	)
}
