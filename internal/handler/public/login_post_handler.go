package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-empolyment/internal/logic/public"
	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"
)

// 登录
func LoginPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewLoginPostLogic(r.Context(), svcCtx)
		resp, err := l.LoginPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
