package router

import (
	"github.com/a406299736/goframe/internal/core"
	"github.com/a406299736/goframe/internal/handler/user_handler"
)

func setDemoRouter(r *resource) {

	back := user_handler.NewUserHandler()

	// 三种验证规则

	// 1. 需要签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.mux.Group("/api/demo", r.middles.Signature())
	{
		//login.GET("/no/detail", back.Detail()) //demo
		login.POST("/no/detail", back.Info()) //demo
	}

	// 2. 需要签名验证、登录验证，无需 RBAC 权限验证
	notRBAC := r.mux.Group("/api/demo", core.WrapAuthHandler(r.middles.Token), r.middles.Signature())
	{
		notRBAC.GET("/not-rabc/detail", back.Info()) // demo
	}

	// 3. 需要签名验证、登录验证、RBAC 权限验证
	api := r.mux.Group("/api/demo", core.WrapAuthHandler(r.middles.Token), r.middles.Signature(), r.middles.RBAC())
	{
		api.POST("/detail", back.Info()) // demo
	}
}
