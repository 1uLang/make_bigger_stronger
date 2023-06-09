package model

import (
	"context"
	"github.com/1uLang/make_bigger_stronger/common/utils/tools"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		ListPage(ctx context.Context, sb squirrel.SelectBuilder, page *tools.Page) ([]*Users, int64, error)
		Count(ctx context.Context, sb squirrel.SelectBuilder) (int64, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

func (m *defaultUsersModel) ListPage(ctx context.Context, sb squirrel.SelectBuilder, page *tools.Page) ([]*Users, int64, error) {
	sb = sb.From(m.table)
	query, args, err := sb.Column("COUNT(1) as count").ToSql()
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = m.QueryRowNoCacheCtx(ctx, &count, query, args...)
	if err != nil {
		return nil, 0, err
	}

	query, args, err = sb.Columns(usersFieldNames...).OrderBy("`created_time` desc").Offset(uint64(page.Offset())).Limit(uint64(page.Size())).ToSql()
	if err != nil {
		return nil, 0, err
	}

	var data []*Users
	err = m.QueryRowsNoCacheCtx(ctx, &data, query, args...)
	if err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

func (m *defaultUsersModel) Count(ctx context.Context, sb squirrel.SelectBuilder) (int64, error) {
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
