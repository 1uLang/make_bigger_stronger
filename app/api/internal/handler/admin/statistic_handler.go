package admin

import (
	"net/http"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/logic/admin"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StatisticHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewStatisticLogic(r.Context(), svcCtx)
		resp, err := l.Statistic()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
