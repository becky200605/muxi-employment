package public

import (
	"context"
	"encoding/base64"
	"net/http"

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

func (l *OriginLogic) Origin(r *http.Request) (resp *types.Response, err error) {
	token := r.Header.Get("Authorization")
		if token == "" {
			resp=&types.Response{
				Code: 403,
				Message: "failed",
				Data: types.OriginResponse{Info: types.Info{Information: "似乎少了些什么呢？再仔细看看前面给出的tips吧"}},
			}
			return
		}
	l1,_:=base64.StdEncoding.DecodeString(token)
	l2,_:=base64.StdEncoding.DecodeString(string(l1))
	rawToken:=string(l2) 
	err=l.svcCtx.JWTHandler.ParseToken(rawToken)
	if err!=nil{
		resp=&types.Response{
				Code: 403,
				Message: "failed",
				Data: types.OriginResponse{Info: types.Info{Information: "似乎不是这个token？再仔细找找吧"}},
			}
			return
	}
	resp=&types.Response{
				Code: 200,
				Message: "failed",
				Data: types.OriginResponse{
					Info: types.Info{
					Information:"我就知道你可以做到！你成功阻止内部背叛，找回 Project-ORIGIN 核心文件。恭喜你成为了择木而犀有限公司的正式员工！",
				},
			},
			}
	return
}
