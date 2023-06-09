package types

import (
	"context"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionLogic {
	return &OptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionLogic) Option() (resp *types.TypeOptionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
