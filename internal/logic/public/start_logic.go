package public

import (
	"context"
	"fmt"

	"muxi-empolyment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开始
func NewStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLogic {
	return &StartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartLogic) Start() (resp string, err error) {
	resp = fmt.Sprintf(`你叫木小犀，一名刚入行的后端开发者。
你怀着热爱入职了一个看似普通的公司--择木而犀有限公司。
而实习期间，主管却悄悄告诉你，公司某款未上线的 AI 产品 “Project-ORIGIN” 于昨日凌晨被内部人士窃取。
你的任务是通过公司内部 API 进行调查并找回被盗数据，以此获得最终宝贵的转正资格。

tips:登录url为 %s/login，需要用户名和密码，你作为实习生还没有权限，但是似乎Mr.Lu 在 %s/debug/log 里面留下了线索...`, l.svcCtx.Config.BASEURL, l.svcCtx.Config.BASEURL)
	return
}
