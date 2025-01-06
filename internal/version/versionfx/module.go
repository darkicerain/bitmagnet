package versionfx

import (
	"bitmagnet-io/bitmagnet/internal/version"
	"bitmagnet-io/bitmagnet/internal/version/healthcheck"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"version",
		fx.Provide(fx.Annotated{
			Name: "app_version",
			Target: func() string {
				v := version.GitTag
				if v == "" {
					v = "unknown"
				}
				return v
			},
		}),
		fx.Provide(healthcheck.New),
	)
}
