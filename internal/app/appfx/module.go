package appfx

import (
  "bitmagnet-io/bitmagnet/internal/app/cmd/classifiercmd"
  "bitmagnet-io/bitmagnet/internal/app/cmd/processcmd"
  "bitmagnet-io/bitmagnet/internal/app/cmd/reprocesscmd"
  "bitmagnet-io/bitmagnet/internal/blocking/blockingfx"
  "bitmagnet-io/bitmagnet/internal/boilerplate/app/boilerplateappfx"
  "bitmagnet-io/bitmagnet/internal/boilerplate/httpserver/httpserverfx"
  "bitmagnet-io/bitmagnet/internal/classifier/classifierfx"
  "bitmagnet-io/bitmagnet/internal/database/databasefx"
  "bitmagnet-io/bitmagnet/internal/database/migrations"
  "bitmagnet-io/bitmagnet/internal/dhtcrawler/dhtcrawlerfx"
  "bitmagnet-io/bitmagnet/internal/downloader/downloaderfx"
  "bitmagnet-io/bitmagnet/internal/gql/gqlfx"
  "bitmagnet-io/bitmagnet/internal/health/healthfx"
  "bitmagnet-io/bitmagnet/internal/importer/importerfx"
  "bitmagnet-io/bitmagnet/internal/metrics/metricsfx"
  "bitmagnet-io/bitmagnet/internal/processor/processorfx"
  "bitmagnet-io/bitmagnet/internal/protocol/dht/dhtfx"
  "bitmagnet-io/bitmagnet/internal/protocol/metainfo/metainfofx"
  "bitmagnet-io/bitmagnet/internal/queue/queuefx"
  "bitmagnet-io/bitmagnet/internal/telemetry/telemetryfx"
  "bitmagnet-io/bitmagnet/internal/tmdb/tmdbfx"
  "bitmagnet-io/bitmagnet/internal/torznab/torznabfx"
  "bitmagnet-io/bitmagnet/internal/version/versionfx"
  "bitmagnet-io/bitmagnet/internal/webui"
  "go.uber.org/fx"
)

func New() fx.Option {
  return fx.Module(
    "app",
    blockingfx.New(),
    boilerplateappfx.New(),
    dhtcrawlerfx.New(),
    dhtfx.New(),
    databasefx.New(),
    gqlfx.New(),
    healthfx.New(),
    httpserverfx.New(),
    importerfx.New(),
    metainfofx.New(),
    metricsfx.New(),
    processorfx.New(),
    queuefx.New(),
    telemetryfx.New(),
    tmdbfx.New(),
    torznabfx.New(),
    versionfx.New(),
    classifierfx.New(),
    downloaderfx.New(),
    // cli commands:
    fx.Provide(
      classifiercmd.New,
      reprocesscmd.New,
      processcmd.New,
    ),
    fx.Provide(webui.New),

    fx.Decorate(migrations.NewDecorator),
  )
}
