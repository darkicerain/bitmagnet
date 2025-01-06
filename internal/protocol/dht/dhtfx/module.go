package dhtfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/protocol"
	"bitmagnet-io/bitmagnet/internal/protocol/dht/client"
	"bitmagnet-io/bitmagnet/internal/protocol/dht/ktable"
	"bitmagnet-io/bitmagnet/internal/protocol/dht/responder"
	"bitmagnet-io/bitmagnet/internal/protocol/dht/server"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"dht",
		configfx.NewConfigModule[server.Config]("dht_server", server.NewDefaultConfig()),
		fx.Provide(
			fx.Annotated{
				Name: "dht_node_id",
				Target: func() protocol.ID {
					return protocol.RandomNodeIDWithClientSuffix()
				},
			},
			client.New,
			ktable.New,
			responder.New,
			server.New,
		),
	)
}
