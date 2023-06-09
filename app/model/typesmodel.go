package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TypesModel = (*customTypesModel)(nil)

type (
	// TypesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTypesModel.
	TypesModel interface {
		typesModel
		Count(ctx context.Context, sb squirrel.SelectBuilder) (int64, error)
	}

	customTypesModel struct {
		*defaultTypesModel
	}
)

// NewTypesModel returns a model for the database table.
func NewTypesModel(conn sqlx.SqlConn, c cache.CacheConf) TypesModel {
	return &customTypesModel{
		defaultTypesModel: newTypesModel(conn, c),
	}
}

func (m *defaultTypesModel) Count(ctx context.Context, sb squirrel.SelectBuilder) (int64, error) {
	sb = sb.From(m.table)
	query, args, err := sb.Column("COUNT(1) as count").ToSql()
	if err != nil {
		return 0, err
	}
	var count int64
	err = m.QueryRowNoCacheCtx(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}
