package private

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-empolyment/internal/logic/private"
	"muxi-empolyment/internal/svc"
)

// 工作
func WorkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := private.NewWorkLogic(r.Context(), svcCtx)
		resp, err := l.Work()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
