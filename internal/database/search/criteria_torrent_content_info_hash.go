package search

import (
	"bitmagnet-io/bitmagnet/internal/database/query"
	"bitmagnet-io/bitmagnet/internal/model"
	"bitmagnet-io/bitmagnet/internal/protocol"
)

func TorrentContentInfoHashCriteria(infoHashes ...protocol.ID) query.Criteria {
	return infoHashCriteria(model.TableNameTorrentContent, infoHashes...)
}
