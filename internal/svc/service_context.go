package svc

import (
	"muxi-empolyment/internal/config"
	"muxi-empolyment/internal/middleware"
	"muxi-empolyment/internal/pkg/ijwt"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	JWTHandler ijwt.JWTHandler
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	JWTHandler:=ijwt.NewJWTHandler(c.Auth.AccessSecret)
	return &ServiceContext{
		Config: c,
		JWTHandler: JWTHandler,
		AuthMiddleware: middleware.NewAuthMiddleware(c,JWTHandler).AuthHandle,
	}
}
