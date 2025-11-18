package public

import (
	"context"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginPostLogic {
	return &LoginPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginPostLogic) LoginPost(req *types.LoginRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
