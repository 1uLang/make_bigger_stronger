package admin

import (
	"context"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/types"
	"github.com/1uLang/make_bigger_stronger/app/model"
	"github.com/1uLang/make_bigger_stronger/common/errors"
	"github.com/1uLang/make_bigger_stronger/common/lib/jwt"
	"github.com/1uLang/make_bigger_stronger/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	set, err := l.svcCtx.SettingsModel.FindOneByKey(l.ctx, model.AdminPasswordPrefix)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.DbError()
	}
	if set == nil { // 获取系统管理员默认密码
		set = &model.Settings{Key: model.AdminPasswordPrefix}
		if l.svcCtx.Config.Setting.DefaultPassword == "" {
			set.Value = utils.MD5(model.DefaultPassword)
		} else {
			set.Value = utils.MD5(l.svcCtx.Config.Setting.DefaultPassword)
		}
		_, _ = l.svcCtx.SettingsModel.Insert(l.ctx, set)
	}
	if req.Account != "admin" || req.Password == set.Value {
		return nil, errors.LoginError()
	}
	formerToken, err := l.svcCtx.JwtInstance.Load(l.ctx, "admin")

	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}
	//	上个token存在 则废弃掉
	if formerToken != "" {
		err = l.svcCtx.JwtInstance.DiscardWithDelete(formerToken)
		//	废弃失败 返回错误
		if err != nil {
			logx.Error(err)
			return nil, errors.SystemError()
		}
	}

	//	生成新的token
	payload := make(jwt.Payload)
	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}

	token, err := l.svcCtx.JwtInstance.Token(payload)
	//	签发新token失败
	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}

	//	存储新的token
	err = l.svcCtx.JwtInstance.Store(l.ctx, "admin", token)
	if err != nil {
		logx.Error(err)
		return nil, errors.SystemError()
	}

	return &types.LoginResp{
		Token: token,
	}, nil
}
