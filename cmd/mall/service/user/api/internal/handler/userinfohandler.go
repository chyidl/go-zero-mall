package handler

import (
	"net/http"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/api/internal/logic"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
