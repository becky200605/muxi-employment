package private

import (
	"context"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OriginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 最初
func NewOriginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OriginLogic {
	return &OriginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OriginLogic) Origin() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
