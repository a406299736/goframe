package router

import (
	"github.com/a406299736/goframe/app/api/controller/demohandler"
	"github.com/a406299736/goframe/app/pkg/core"
)

func setDemo2Router2(r *resource) {

	demo2 := demohandler.New2(r.db, r.cache)

	// 三种验证规则

	// 1. 需要签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.mux.Group("/api/demo2", core.WrapAuthHandler(r.middles.Token))
	{
		//login.GET("/no/detail", back.Detail()) //demo
		login.GET("/no/detail", demo2.Info()) //demo
		login.POST("/no/c", demo2.Create())   //demo
	}

	// 2. 需要签名验证、登录验证，无需 RBAC 权限验证
	//notRBAC := r.mux.Group("/api/demo", core.WrapAuthHandler(r.middles.Token), r.middles.Signature())
	{
		//notRBAC.GET("/not-rabc/detail", back.Detail()) // demo
	}

	// 3. 需要签名验证、登录验证、RBAC 权限验证
	//api := r.mux.Group("/api/demo", core.WrapAuthHandler(r.middles.Token), r.middles.Signature(), r.middles.RBAC())
	{
		//api.POST("/detail", back.Detail()) // demo
	}
}
