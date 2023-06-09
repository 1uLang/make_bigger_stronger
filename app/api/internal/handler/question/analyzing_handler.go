package question

import (
	"net/http"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/logic/question"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AnalyzingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnalyzingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := question.NewAnalyzingLogic(r.Context(), svcCtx)
		resp, err := l.Analyzing(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
