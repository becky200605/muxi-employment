package public

import (
	"context"
	"fmt"

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
	resp=&types.Response{}
	logInfo:=[]string{
		"[INFO] 2025-11-20 09:21:56 用户 'root' 登录失败",
		"[ERROR] 无法读取 image: 72.jpg",
		"[INFO] npc-Lu 成功通关《丝之歌》第一结局，用时 72 小时",
		"[WARN] base64-cult: new believer registered",
		fmt.Sprintf("[INFO] user=alpha debug-mode=true password=%s",l.svcCtx.Config.User.EncodedPassword),
		"[INFO] 内部虚拟助手 “Archives-72” 请求查看通关纪念图",
	}
	resp.Data=types.LogResponse{
		Info: types.Info{
           Information: "似乎是某个开发人员上线时忘记关闭的调试日志接口，里面好像记录了什么重要的东西呢",
        },
		Logs: logInfo,
	}
	return
}
