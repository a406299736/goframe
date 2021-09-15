### Usage

#### 部署时设置环境变量 export PROJECT_PATH=项目绝对路径.
        使用goland IDE可编辑configurations中environment项，设置为PROJECT_PATH=项目绝对路径

#### 1. app/router/router.go 注册路由或注册路由分组.

#### 2. 一些中间件在 app/router/middleware/ 下注册实现.

#### 3. 编写路由规则和鉴权规则.
            参考 app/router/router-demo.go
            推荐简化版示例: app/router/router-demo2.go

#### 4. 实现 service 侧接口时 统一返回 errors.Er类型错误；handler 侧接收到 errors.Er 错误时，调用用context.Failed(errno.Error)压入 _abort_error_ 调用栈，core包最后统一处理日志.
##### ①. 参考 
            app/router/router_demo.go
            app/api/controller/damo-handler/handler.go
            app/api/service/demo/service.go
            推荐简化版示例:
            app/router/router_demo2.go
            app/api/controller/damo-handler/handler2.go
            app/api/service/demo/service2.go

#### 5. 监控 
            {{host}}/debug/pprof/
            {{host}}/metrics
            {{host}}/system/health

#### 6. 常用请求
            ctx.ShouldBindJSON(req) // 获取body值 json格式
            ctx.ShouldBindURI(req)  // 获取url参数 
            ... ...
#### 7. CRUD
            简化版示例: app/api/service/demo/service2.go

            dc := demo.NewDemo()
            dc.Aspirations = "demo"
            id, err := dc.Create(src.db.GetDbR().WithContext(c.RequestContext()))
            

            dm := demo.NewDemoQueryBuilder()
	        all, _ := dm.WhereIdIn([]int32{1, 2, 3}).
		    WhereMobileNotIn([]string{"111", "222"}).
		    OrderById(false)
		    list := dm.QueryAll(src.db.GetDbR().WithContext(c.RequestContext())))
            
            err := dm.Updates(src.db.GetDbR().WithContext(c.RequestContext())))

            err := dm.Delete(src.db.GetDbR().WithContext(c.RequestContext())))

#### 8. 代码生成工具，shell脚本 参数：库名 表名；默认连接为测试库，可在cmd/mysqlmd/main.go 修改，或加入flag参数列表.
             ./scripts/gormgen.sh wm_class wm_user

#### 9. rocketMQ
            参考 console/cmd/rocketMQ.go
            use r := mq.New(nil)
            r.Producer(...).Push(...) // 生产
            or r.Consumer(...).pull(doFunc) // 消费

#### 10. console 如果使用Context接口,需要使用 core.NewCmdContext(logger)
            参考 console/cmd/mockDemo.go
            go run console/main.go MockDemo 
            or go build -o cmd
            ./cmd MockDemo