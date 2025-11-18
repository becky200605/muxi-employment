package private

import (
	"context"

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

func (l *AssetLogic) Asset() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
