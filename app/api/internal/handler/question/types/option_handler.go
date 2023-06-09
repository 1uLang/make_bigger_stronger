package types

import (
	"net/http"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/logic/question/types"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := types.NewOptionLogic(r.Context(), svcCtx)
		resp, err := l.Option()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
