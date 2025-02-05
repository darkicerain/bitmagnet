package gqlfx

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/lazy"
	"bitmagnet-io/bitmagnet/internal/boilerplate/worker"
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/database/search"
	"bitmagnet-io/bitmagnet/internal/gql"
	"bitmagnet-io/bitmagnet/internal/gql/config"
	"bitmagnet-io/bitmagnet/internal/gql/httpserver"
	"bitmagnet-io/bitmagnet/internal/gql/resolvers"
	"bitmagnet-io/bitmagnet/internal/health"
	"bitmagnet-io/bitmagnet/internal/metrics/queuemetrics"
	"bitmagnet-io/bitmagnet/internal/metrics/torrentmetrics"
	"bitmagnet-io/bitmagnet/internal/queue/manager"
	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"graphql",
		fx.Provide(
			config.New,
			httpserver.New,
			func(
				lcfg lazy.Lazy[gql.Config],
			) lazy.Lazy[graphql.ExecutableSchema] {
				return lazy.New(func() (graphql.ExecutableSchema, error) {
					cfg, err := lcfg.Get()
					if err != nil {
						return nil, err
					}
					return gql.NewExecutableSchema(cfg), nil
				})
			},
		),
		fx.Provide(
			func(p Params) Result {
				return Result{
					Resolver: lazy.New(func() (*resolvers.Resolver, error) {
						ch, err := p.Checker.Get()
						if err != nil {
							return nil, err
						}
						s, err := p.Search.Get()
						if err != nil {
							return nil, err
						}
						d, err := p.Dao.Get()
						if err != nil {
							return nil, err
						}
						qmc, err := p.QueueMetricsClient.Get()
						if err != nil {
							return nil, err
						}
						qm, err := p.QueueManager.Get()
						if err != nil {
							return nil, err
						}
						tm, err := p.TorrentMetricsClient.Get()
						if err != nil {
							return nil, err
						}
						return &resolvers.Resolver{
							Dao:                  d,
							Search:               s,
							Checker:              ch,
							QueueMetricsClient:   qmc,
							QueueManager:         qm,
							TorrentMetricsClient: tm,
						}, nil
					}),
				}
			},
		),
		// inject resolver dependencies avoiding a circular dependency:
		fx.Invoke(func(
			resolver lazy.Lazy[*resolvers.Resolver],
			workers worker.Registry,
		) {
			resolver.Decorate(func(r *resolvers.Resolver) (*resolvers.Resolver, error) {
				r.Workers = workers
				return r, nil
			})
		}),
	)
}

type Params struct {
	fx.In
	Search               lazy.Lazy[search.Search]
	Dao                  lazy.Lazy[*dao.Query]
	Checker              lazy.Lazy[health.Checker]
	QueueMetricsClient   lazy.Lazy[queuemetrics.Client]
	QueueManager         lazy.Lazy[manager.Manager]
	TorrentMetricsClient lazy.Lazy[torrentmetrics.Client]
}

type Result struct {
	fx.Out
	Resolver lazy.Lazy[*resolvers.Resolver]
}
