package torznab

import (
	"bitmagnet-io/bitmagnet/internal/model"
)

type SearchRequest struct {
	Query    string
	Type     string
	Cats     []int
	ImdbId   model.NullString
	TmdbId   model.NullString
	Season   model.NullInt
	Episode  model.NullInt
	Attrs    []string
	Extended bool
	Limit    model.NullUint
	Offset   model.NullUint
}
