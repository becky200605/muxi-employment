package private

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"os"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取图片
func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetLogic) Asset(w http.ResponseWriter, r *http.Request, req *types.AssetPathReq) error {
	switch req.ImageId {
	case "72":
		return l.Assert72(w, "./internal/data/72.png")
	default:
		httpx.OkJsonCtx(r.Context(), w, &types.Response{
			Code:    200,
			Message: "success",
			Data:    &types.AssetResponse{Info: types.Info{Information: "嘿嘿，暂时没有这个图片哦"}},
		})
		return nil
	}
}

func (l *AssetLogic) Assert72(w http.ResponseWriter, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置 Content-Type
	w.Header().Set("Content-Type", "image/png")

	// 读取文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// 直接写入响应
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	return err
}
