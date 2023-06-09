package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionsModel = (*customQuestionsModel)(nil)

type (
	// QuestionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionsModel.
	QuestionsModel interface {
		questionsModel
		Count(ctx context.Context, sb squirrel.SelectBuilder) (int64, error)
	}

	customQuestionsModel struct {
		*defaultQuestionsModel
	}
)

// NewQuestionsModel returns a model for the database table.
func NewQuestionsModel(conn sqlx.SqlConn, c cache.CacheConf) QuestionsModel {
	return &customQuestionsModel{
		defaultQuestionsModel: newQuestionsModel(conn, c),
	}
}

func (m *defaultQuestionsModel) Count(ctx context.Context, sb squirrel.SelectBuilder) (int64, error) {
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
