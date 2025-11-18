package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-empolyment/internal/logic/public"
	"muxi-empolyment/internal/svc"
)

// 获取日志
func LogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewLogLogic(r.Context(), svcCtx)
		resp, err := l.Log()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
