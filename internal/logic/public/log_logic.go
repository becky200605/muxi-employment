package public

import (
	"context"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取日志
func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLogic {
	return &LogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogLogic) Log() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
