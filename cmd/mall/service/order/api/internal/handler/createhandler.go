package handler

import (
	"net/http"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/logic"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
