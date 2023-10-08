package router

import (
	"github.com/a406299736/goframe/app/api/controller/democtrl"
	"github.com/a406299736/goframe/app/pkg/core"
	"io"
	"os"
)

func setDemo2Router2(r *resource) {

	demo2 := democtrl.New2()

	// 访问路由(/filename)，直接下载文件
	r.mux.Group("/").GET("filename", func(c core.Context) {
		filepath := ""
		file, _ := os.OpenFile(filepath, os.O_RDONLY, 0666)
		bytes, _ := io.ReadAll(file)
		c.GinCtx().Data(200, "application/octet-stream", bytes)
		c.SetForceBreak(true)
	})

	// 访问路由(/filename)，浏览文件
	r.mux.Group("/").GET("filename1", func(c core.Context) {
		filepath := ""
		c.GinCtx().File(filepath)
		c.SetForceBreak(true)
	})

	// 三种验证规则

	// 1. 无需签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.mux.Group("/api/demo2")
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
