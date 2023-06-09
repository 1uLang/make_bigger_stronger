package question

import (
	"context"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateLogic {
	return &TemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateLogic) Template() error {
	// todo: add your logic here and delete this line

	return nil
}
