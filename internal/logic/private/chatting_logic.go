package private

import (
	"context"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChattingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 聊天
func NewChattingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChattingLogic {
	return &ChattingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChattingLogic) Chatting() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
