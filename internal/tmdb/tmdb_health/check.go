package tmdb_health

import (
	"bitmagnet-io/bitmagnet/internal/boilerplate/lazy"
	"bitmagnet-io/bitmagnet/internal/health"
	"bitmagnet-io/bitmagnet/internal/tmdb"
	"context"
	"time"
)

func NewCheck(
	enabled bool,
	client lazy.Lazy[tmdb.Client],
) health.Check {
	return health.Check{
		Name:    "tmdb",
		Timeout: time.Second * 30,
		IsActive: func() bool {
			return enabled
		},
		Check: func(ctx context.Context) error {
			c, err := client.Get()
			if err != nil {
				return err
			}
			return c.ValidateApiKey(ctx)
		},
	}
}
