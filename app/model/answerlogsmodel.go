package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnswerLogsModel = (*customAnswerLogsModel)(nil)

type (
	// AnswerLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnswerLogsModel.
	AnswerLogsModel interface {
		answerLogsModel
	}

	customAnswerLogsModel struct {
		*defaultAnswerLogsModel
	}
)

// NewAnswerLogsModel returns a model for the database table.
func NewAnswerLogsModel(conn sqlx.SqlConn, c cache.CacheConf) AnswerLogsModel {
	return &customAnswerLogsModel{
		defaultAnswerLogsModel: newAnswerLogsModel(conn, c),
	}
}

func (m *defaultAnswerLogsModel) AnswersCount(ctx context.Context, sb squirrel.SelectBuilder) (int64, error) {
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
