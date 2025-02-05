package search

import (
	"bitmagnet-io/bitmagnet/internal/database/dao"
	"bitmagnet-io/bitmagnet/internal/database/query"
	"bitmagnet-io/bitmagnet/internal/maps"
	"bitmagnet-io/bitmagnet/internal/model"
	"database/sql/driver"
	"gorm.io/gen/field"
)

var VideoResolutionCriteria = torrentContentAttributeCriteria[model.VideoResolution](videoResolutionField)

var Video3DCriteria = torrentContentAttributeCriteria[model.Video3D](video3dField)

func torrentContentAttributeCriteria[T attribute](getFld func(*dao.Query) field.Field) func(...T) query.Criteria {
	return func(values ...T) query.Criteria {
		return query.DaoCriteria{
			Conditions: func(ctx query.DbContext) ([]field.Expr, error) {
				q := ctx.Query()
				fld := getFld(q)
				valuers := make([]driver.Valuer, 0, len(values))
				for _, v := range values {
					valuers = append(valuers, v)
				}
				return []field.Expr{fld.In(valuers...)}, nil
			},
			Joins: maps.NewInsertMap(
				maps.MapEntry[string, struct{}]{Key: model.TableNameTorrentContent},
			),
		}
	}
}
