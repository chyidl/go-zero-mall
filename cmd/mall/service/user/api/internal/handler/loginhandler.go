package handler

import (
	"net/http"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/api/internal/logic"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/api/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
