package search

import (
	"bitmagnet-io/bitmagnet/internal/database/query"
	"bitmagnet-io/bitmagnet/internal/model"
)

const ContentGenreFacetKey = "content_genre"

func TorrentContentGenreFacet(options ...query.FacetOption) query.Facet {
	return torrentContentCollectionFacet{
		FacetConfig: query.NewFacetConfig(
			append([]query.FacetOption{
				query.FacetHasKey(ContentGenreFacetKey),
				query.FacetHasLabel("Genre"),
				query.FacetUsesLogic(model.FacetLogicAnd),
				query.FacetHasAggregationOption(query.RequireJoin(model.TableNameTorrentContent)),
				query.FacetTriggersCte(),
			}, options...)...,
		),
		collectionType: "genre",
	}
}
