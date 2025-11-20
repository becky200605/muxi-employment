package middleware

import (
	"fmt"
	"muxi-empolyment/internal/config"
	"muxi-empolyment/internal/pkg/ijwt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	Cfg config.Config
	r   ijwt.JWTHandler
}

func NewAuthMiddleware(cfg config.Config, r ijwt.JWTHandler) *AuthMiddleware {
	return &AuthMiddleware{
		Cfg: cfg,
		r:   r,
	}
}

func (m *AuthMiddleware) AuthHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid authoriztion"))
			return
		}
		err := m.r.ParseToken(token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("解析token失败:%v", err))
			return
		}
		next(w, r)
	}
}
