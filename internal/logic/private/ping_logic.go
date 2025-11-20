package private

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Ping
func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type IPInfo struct {
	Count      int
	FirstVisit time.Time
}

var (
	ipMap         sync.Map
	requestLimit  = 5
	limitDuration = 3 * time.Second
)

func (l *PingLogic) Ping(r *http.Request) (resp *types.Response, err error) {
	ip := getRealIP(r)
	value, _ := ipMap.LoadOrStore(ip, &IPInfo{
		Count:      0,
		FirstVisit: time.Now(),
	})
	ipInfo := value.(*IPInfo)
	ipInfo.Count++
	if ipInfo.Count >= requestLimit && time.Since(ipInfo.FirstVisit) <= limitDuration {
		resp = &types.Response{
			Code:    200,
			Message: "success",
			Data: types.PingResponse{
				Info: types.Info{
					Information: "刷卡成功！你有权限访问 /origin 了",
				},
				Tips: types.Tips{
					Tips: "Mr.Lu有些神秘要求，他希望你请求头中的token并非 jwt,而是 base64(base64(token))",
				},
			},
		}
		return
	}

	if time.Since(ipInfo.FirstVisit) > limitDuration {
		ipInfo.Count = 1
		ipInfo.FirstVisit = time.Now()
	}
	resp = &types.Response{
		Code:    403,
		Message: "failed",
		Data: types.PingResponse{
			Info: types.Info{
				Information: "访问太慢，检测到异常人员，已上报。",
			},
			Tips: types.Tips{
				Tips: "想象一下保安为什么都能认出业主，因为他们不屑一顾（bushi）。一次一次请求太慢了，过于温柔了，是时候抛弃礼貌了 --来自Mr.Lu",
			},
		},
	}
	return
}

func getRealIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
