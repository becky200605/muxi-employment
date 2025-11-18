package public

import (
	"context"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginGetLogic {
	return &LoginGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginGetLogic) LoginGet() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
