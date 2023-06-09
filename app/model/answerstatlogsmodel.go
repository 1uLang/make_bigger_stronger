package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnswerStatLogsModel = (*customAnswerStatLogsModel)(nil)

type (
	// AnswerStatLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnswerStatLogsModel.
	AnswerStatLogsModel interface {
		answerStatLogsModel
	}

	customAnswerStatLogsModel struct {
		*defaultAnswerStatLogsModel
	}
)

// NewAnswerStatLogsModel returns a model for the database table.
func NewAnswerStatLogsModel(conn sqlx.SqlConn, c cache.CacheConf) AnswerStatLogsModel {
	return &customAnswerStatLogsModel{
		defaultAnswerStatLogsModel: newAnswerStatLogsModel(conn, c),
	}
}
