package search

import (
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/database/query"
	"bitmagnet-io/bitmagnet/internal/model"
	"context"
)

type TorrentFilesResult = query.GenericResult[model.TorrentFile]

type TorrentFilesSearch interface {
	TorrentFiles(ctx context.Context, options ...query.Option) (TorrentFilesResult, error)
}

func (s search) TorrentFiles(ctx context.Context, options ...query.Option) (TorrentFilesResult, error) {
	return query.GenericQuery[model.TorrentFile](
		ctx,
		s.q,
		query.Options(append([]query.Option{query.SelectAll()}, options...)...),
		model.TableNameTorrentFile,
		func(ctx context.Context, q *dao.Query) query.SubQuery {
			return query.GenericSubQuery[dao.ITorrentFileDo]{
				SubQuery: q.TorrentFile.WithContext(ctx).ReadDB(),
			}
		},
	)
}
