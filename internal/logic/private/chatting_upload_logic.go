package private

import (
	"bytes"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChattingUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传图片
func NewChattingUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChattingUploadLogic {
	return &ChattingUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *ChattingUploadLogic) ChattingUpload(w http.ResponseWriter, r *http.Request, req *types.ChattingUploadRequest) error {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    400,
			Message: "Content-Type 必须为 multipart/form-data",
		})
		return nil
	}

	err := r.ParseMultipartForm(322 * 1024 * 104)
	if err != nil {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    400,
			Message: "解析表单失败: " + err.Error(),
		})
		return nil
	}
	defer r.MultipartForm.RemoveAll()

	fileHeaders := r.MultipartForm.File["file"]
	if len(fileHeaders) == 0 {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    400,
			Message: "缺少 file 字段",
		})
		return nil
	}

	fileHeader := fileHeaders[0]
	if fileHeader.Size == 0 {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    400,
			Message: "上传的文件为空",
		})
		return nil
	}

	srcFile, err := fileHeader.Open()
	if err != nil {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    500,
			Message: "无法打开上传文件",
		})
		return nil
	}
	defer srcFile.Close()

	uploadedData, err := io.ReadAll(srcFile)
	if err != nil {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    500,
			Message: "读取上传文件失败",
		})
		return nil
	}

	existingFilePath := "./internal/data/72.png"
	existingData, err := os.ReadFile(existingFilePath)
	if err != nil {
		httpx.OkJsonCtx(ctx, w, &types.Response{
			Code:    500,
			Message: "读取目标文件失败: " + err.Error(),
		})
		return nil
	}

	same := bytes.Equal(uploadedData, existingData)

	if same {
		httpx.OkJsonCtx(r.Context(), w, &types.Response{
			Code:    200,
			Message: "success",
			Data: &types.ChattingResponse{
				Info: types.Info{Information: "\"不错嘛新人！你跟我一样，都是丝之歌信徒。\\n我听说Project-ORIGIN 的核心文件被内部人加密后上传到 /origin，但你需要正确的 token 形式 \\n 对于这个，我建议你可以去找 npc-security 聊聊\""},
			},
		})
	} else {
		httpx.OkJsonCtx(r.Context(), w, &types.Response{
			Code:    200,
			Message: "success",
			Data: &types.ChattingResponse{
				Info: types.Info{Information: "谢谢你的图片，但我并不对其感兴趣"},
			},
		})
	}
	return nil
}
