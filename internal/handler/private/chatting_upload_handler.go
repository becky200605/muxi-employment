package private

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-empolyment/internal/logic/private"
	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"
)

// 上传图片
func ChattingUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChattingUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := private.NewChattingUploadLogic(r.Context(), svcCtx)
		err := l.ChattingUpload(w, r, &req)
		if err != nil {
			logx.Errorf("ChattingUploadHandler error: %v", err)
		}
	}
}
