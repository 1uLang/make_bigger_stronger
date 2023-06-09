// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBiggerStrongerUsersIdPrefix     = "cache:biggerStronger:users:id:"
	cacheBiggerStrongerUsersOpenIdPrefix = "cache:biggerStronger:users:openId:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByOpenId(ctx context.Context, openId string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id         int64  `db:"id"`
		OpenId     string `db:"openId"`      // 用户微信open ID
		Name       string `db:"name"`        // 用户名称 默认微信名称
		Count      int64  `db:"count"`       // 答题次数
		Stat       string `db:"stat"`        // 答题统计：答对数/总数
		CreateTime int64  `db:"create_time"` // 创建时间
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	biggerStrongerUsersIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersIdPrefix, data.Id)
	biggerStrongerUsersOpenIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersOpenIdPrefix, data.OpenId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.OpenId, data.Name, data.Count, data.Stat)
	}, biggerStrongerUsersIdKey, biggerStrongerUsersOpenIdKey)
	return ret, err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	biggerStrongerUsersIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, biggerStrongerUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByOpenId(ctx context.Context, openId string) (*Users, error) {
	biggerStrongerUsersOpenIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersOpenIdPrefix, openId)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, biggerStrongerUsersOpenIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `openId` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, openId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	biggerStrongerUsersIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersIdPrefix, data.Id)
	biggerStrongerUsersOpenIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersOpenIdPrefix, data.OpenId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.OpenId, data.Name, data.Count, data.Stat, data.Id)
	}, biggerStrongerUsersIdKey, biggerStrongerUsersOpenIdKey)
	return err
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	biggerStrongerUsersIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersIdPrefix, id)
	biggerStrongerUsersOpenIdKey := fmt.Sprintf("%s%v", cacheBiggerStrongerUsersOpenIdPrefix, data.OpenId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, biggerStrongerUsersIdKey, biggerStrongerUsersOpenIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBiggerStrongerUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
