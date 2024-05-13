package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"v1/api/internal/logic"
	"v1/api/internal/svc"
	"v1/api/internal/types"
)

func ShortenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), svcCtx)
		resp, err := l.Shorten(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
