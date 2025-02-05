package server

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/lazy"
	"bitmagnet-io/bitmagnet/internal/boilerplate/worker"
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/queue/handler"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type Params struct {
	fx.In
	Query lazy.Lazy[*dao.Query]
	//PgxPool  lazy.Lazy[*pgxpool.Pool]
	Handlers []lazy.Lazy[handler.Handler] `group:"queue_handlers"`
	Logger   *zap.SugaredLogger
}

type Result struct {
	fx.Out
	Worker worker.Worker `group:"workers"`
}

func New(p Params) Result {
	stopped := make(chan struct{})
	return Result{
		Worker: worker.NewWorker(
			"queue_server",
			fx.Hook{
				OnStart: func(ctx context.Context) error {
					//pool, err := p.PgxPool.Get()
					//if err != nil {
					//	return err
					//}
					query, err := p.Query.Get()
					if err != nil {
						return err
					}
					handlers := make([]handler.Handler, 0, len(p.Handlers))
					for _, lh := range p.Handlers {
						h, err := lh.Get()
						if err != nil {
							return err
						}
						handlers = append(handlers, h)
					}
					srv := server{
						stopped: stopped,
						query:   query,
						//pool:       pool,
						handlers:   handlers,
						gcInterval: time.Minute * 10,
						logger:     p.Logger.Named("queue"),
					}
					return srv.Start(context.Background())
				},
				OnStop: func(ctx context.Context) error {
					close(stopped)
					return nil
				},
			},
		),
	}
}
