package question

import (
	"context"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalyzingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnalyzingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzingLogic {
	return &AnalyzingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnalyzingLogic) Analyzing(req *types.AnalyzingReq) (resp *types.AnalyzingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
