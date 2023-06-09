package svc

import (
	"github.com/1uLang/make_bigger_stronger/app/api/internal/config"
	"github.com/1uLang/make_bigger_stronger/app/model"
	"github.com/1uLang/make_bigger_stronger/common/lib/jwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Redis
	UserModel       model.UsersModel
	SettingsModel   model.SettingsModel
	QuestionsModel  model.QuestionsModel
	TypesModel      model.TypesModel
	AnswerLogsModel model.AnswerLogsModel
	JwtInstance     jwt.JWT
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.MustNewRedis(c.RedisConf)
	srvCtx := &ServiceContext{
		Config:          c,
		RedisClient:     redisClient,
		UserModel:       model.NewUsersModel(conn, c.Mysql.DbCache),
		SettingsModel:   model.NewSettingsModel(conn, c.Mysql.DbCache),
		QuestionsModel:  model.NewQuestionsModel(conn, c.Mysql.DbCache),
		TypesModel:      model.NewTypesModel(conn, c.Mysql.DbCache),
		AnswerLogsModel: model.NewAnswerLogsModel(conn, c.Mysql.DbCache),
	}
	srvCtx.JwtInstance = jwt.NewJWT(c.AuthCfg.AccessSecret, c.AuthCfg.AccessExpire, jwt.SetRedis(redisClient), jwt.SetBlackListOpt(true))
	return srvCtx
}
