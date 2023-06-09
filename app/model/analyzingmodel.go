package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnalyzingModel = (*customAnalyzingModel)(nil)

type (
	// AnalyzingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnalyzingModel.
	AnalyzingModel interface {
		analyzingModel
	}

	customAnalyzingModel struct {
		*defaultAnalyzingModel
	}
)

// NewAnalyzingModel returns a model for the database table.
func NewAnalyzingModel(conn sqlx.SqlConn, c cache.CacheConf) AnalyzingModel {
	return &customAnalyzingModel{
		defaultAnalyzingModel: newAnalyzingModel(conn, c),
	}
}
