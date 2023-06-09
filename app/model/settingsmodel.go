package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SettingsModel = (*customSettingsModel)(nil)

const AdminPasswordPrefix = "adminPassword"
const DefaultPassword = "123456"

type (
	// SettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSettingsModel.
	SettingsModel interface {
		settingsModel
	}

	customSettingsModel struct {
		*defaultSettingsModel
	}
)

// NewSettingsModel returns a model for the database table.
func NewSettingsModel(conn sqlx.SqlConn, c cache.CacheConf) SettingsModel {
	return &customSettingsModel{
		defaultSettingsModel: newSettingsModel(conn, c),
	}
}
