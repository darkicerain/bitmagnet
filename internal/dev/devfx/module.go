package devfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/boilerplatefx"
	"bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"bitmagnet-io/bitmagnet/internal/database"
	"bitmagnet-io/bitmagnet/internal/database/migrations"
	"bitmagnet-io/bitmagnet/internal/database/postgres"
	"bitmagnet-io/bitmagnet/internal/dev/app/cmd/gormcmd"
	"bitmagnet-io/bitmagnet/internal/dev/app/cmd/migratecmd"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"dev",
		boilerplatefx.New(),
		configfx.NewConfigModule[postgres.Config]("postgres", postgres.NewDefaultConfig()),
		fx.Provide(database.New),
		fx.Provide(migrations.New),
		fx.Provide(postgres.New),
		fx.Provide(gormcmd.New),
		fx.Provide(migratecmd.New),
	)
}
