package private

import (
	"context"
	"math/rand"
	"time"

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

func (l *ChattingLogic) Chatting(req *types.ChattingPathReq) (resp *types.Response, err error) {
	switch req.NpcName {
	case "":
		return Empty()
	case "npc-lu":
		return NpcLu()
	case "npc-security":
		return NpcSecurity()
	default:
		return Default(req.NpcName)
	}
}

func Empty() (resp *types.Response, err error) {
	return &types.Response{
		Code:    200,
		Message: "success",
		Data: &types.ChattingResponse{
			Info: types.Info{Information: "吼吼哈嘿，快使用 get 请求 /chatting/<npc-name> 聊天吧"},
			Tips: types.Tips{Tips: "不知道和谁聊天，不如回忆一下 Mr.Lu 的后门中的内容吧"},
		},
	}, nil
}

// NpcLu npc-lu 聊天，tips 随机
func NpcLu() (resp *types.Response, err error) {
	tipsOptions := []string{
		"npc-lu 非常的忙碌，但是他很高兴看到一张 丝之歌的通关图片，最好是72小时通关的，这样他也许会告诉一些东西，也许你可以回忆一下日志的内容，顺便去 /asset/<image-id> 看看呢",
		"也许你可以使用 post 方式将你获得的图片给 npc-lu 呢，请求体是 {\"image\":\"xxxx\"}",
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	tips := tipsOptions[rand.Intn(len(tipsOptions))]

	return &types.Response{
		Code:    200,
		Message: "success",
		Data: &types.ChattingResponse{
			Info: types.Info{Information: "你和 npc-lu 聊天，感受到了他的忙碌，多聊天几次试试呢，也许会告诉你一些东西"},
			Tips: types.Tips{Tips: tips},
		},
	}, nil
}

func NpcSecurity() (resp *types.Response, err error) {
	return &types.Response{
		Code:    200,
		Message: "success",
		Data: &types.ChattingResponse{
			Info: types.Info{"请前往 /ping 刷卡"},
		},
	}, nil
}

func Default(name string) (resp *types.Response, err error) {
	return &types.Response{
		Code:    200,
		Message: "success",
		Data: &types.ChattingResponse{
			Info: types.Info{Information: "你和 " + name + " 聊天，但他好像很神秘..."},
			Tips: types.Tips{Tips: "也许你可以尝试其他 NPC 的互动方式"},
		},
	}, nil
}
