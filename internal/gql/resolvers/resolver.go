package resolvers

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/worker"
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/database/search"
	"bitmagnet-io/bitmagnet/internal/health"
	"bitmagnet-io/bitmagnet/internal/metrics/queuemetrics"
	"bitmagnet-io/bitmagnet/internal/metrics/torrentmetrics"
	"bitmagnet-io/bitmagnet/internal/queue/manager"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Dao                  *dao.Query
	Search               search.Search
	Workers              worker.Registry
	Checker              health.Checker
	QueueMetricsClient   queuemetrics.Client
	QueueManager         manager.Manager
	TorrentMetricsClient torrentmetrics.Client
}
