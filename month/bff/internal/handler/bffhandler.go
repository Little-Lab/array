package handler

import (
	"net/http"

	"bff/internal/logic"
	"bff/internal/svc"
	"bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func bffHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BFFRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBffLogic(r.Context(), svcCtx)
		resp, err := l.Bff(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
