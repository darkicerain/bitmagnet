package metainfofx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/protocol/metainfo/banning"
	"bitmagnet-io/bitmagnet/internal/protocol/metainfo/metainforequester"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"metainfo",
		configfx.NewConfigModule[metainforequester.Config]("metainfo_requester", metainforequester.NewDefaultConfig()),
		fx.Provide(
			metainforequester.New,
			banning.New,
		),
	)
}
