package types

import (
	"net/http"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/logic/question/types"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TypeEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := types.NewEditLogic(r.Context(), svcCtx)
		err := l.Edit(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
