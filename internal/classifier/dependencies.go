package classifier

import (
	"bitmagnet-io/bitmagnet/internal/tmdb"
)

type dependencies struct {
	search     LocalSearch
	tmdbClient tmdb.Client
}
