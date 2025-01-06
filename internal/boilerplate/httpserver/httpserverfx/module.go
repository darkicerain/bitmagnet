package httpserverfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/boilerplate/httpserver"
	"bitmagnet-io/bitmagnet/internal/boilerplate/httpserver/cors"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"http_server",
		configfx.NewConfigModule[httpserver.Config]("http_server", httpserver.NewDefaultConfig()),
		fx.Provide(
			httpserver.New,
			cors.New,
		),
	)
}
