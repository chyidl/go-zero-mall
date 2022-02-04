package handler

import (
	"net/http"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/api/internal/logic"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/api/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
