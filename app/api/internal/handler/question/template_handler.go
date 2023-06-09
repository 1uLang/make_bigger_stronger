package question

import (
	"net/http"

	"github.com/1uLang/make_bigger_stronger/app/api/internal/logic/question"
	"github.com/1uLang/make_bigger_stronger/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := question.NewTemplateLogic(r.Context(), svcCtx)
		err := l.Template()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
