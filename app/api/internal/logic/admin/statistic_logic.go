package admin

import (
	"context"
	"github.com/1uLang/make_bigger_stronger/common/errors"
	"github.com/Masterminds/squirrel"
	"time"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticLogic {
	return &StatisticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticLogic) Statistic() (resp *types.StatisticResp, err error) {
	// 统计用户数
	users, err := l.svcCtx.UserModel.Count(l.ctx, squirrel.Select())
	if err != nil {
		return nil, errors.DbError()
	}
	// 统计题库数
	questions, err := l.svcCtx.QuestionsModel.Count(l.ctx, squirrel.Select())
	if err != nil {
		return nil, errors.DbError()
	}
	// 统计题库数
	tps, err := l.svcCtx.TypesModel.Count(l.ctx, squirrel.Select())
	if err != nil {
		return nil, errors.DbError()
	}

	// 统计题库数
	start, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	todayAnswers, err := l.svcCtx.AnswerLogsModel.AnswersCount(l.ctx, squirrel.Select().Where("`create_time` <= ? and `create_time` >= ?", start.Unix(), start.Unix()+24*2600))
	if err != nil {
		return nil, errors.DbError()
	}
	return &types.StatisticResp{Users: users, Questions: questions, Types: tps, TodayAnswer: todayAnswers}, nil
}
