package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-empolyment/internal/logic/public"
	"muxi-empolyment/internal/svc"
)

// 开始
func StartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewStartLogic(r.Context(), svcCtx)
		resp, err := l.Start()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} 
		w.Header().Set("Content-Type","text/plain; charset=utf-8")
		w.Write([]byte(resp))
	}
}
