### Usage
wiki地址: http://wiki.weimiaocaishang.com/pages/viewpage.action?pageId=29118938

#### 部署时设置环境变量 export PROJECT_PATH=项目绝对路径.
        使用goland IDE可编辑configurations中environment项，设置为PROJECT_PATH=项目绝对路径
        go build main.go
         ./main -env test

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
            ctx.ShouldBindForm(req)  // 获取url参数 定义结构体tag: `form:"xxx"`
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

#### 11. 新增支持apollo, 使用简单:
            简述注意事项: 配置项在 ./configs/xxx_configs.tmol, 新增后在 ./configs/configs.go
            新增viper结构体; 建议namespaceName=默认(application), cluster=默认(default),
            这样获取conf时,可以在New函数内不用传参, 如果需要连接非默认配置,
            则在New时需要传入WithConfig() 或 WithNamespace();
            使用如下:
            conf, err := apollo.New(apollo.WithNamespace("application"))
	        if err != nil {
		        fmt.Printf("%+v", err)
	        }
            // 获取 apollo配置key为USER_LIST的值,若key不存在则返回默认值
	        fmt.Println("apollo value USER_LIST:", conf.GetStringValue("USER_LIST", "2222"))
            

常见问题:
1. 运行测试用例时,如果出现找不到配置文件等类似错误, 可以手动执行:
export PROJECT_PATH=项目绝对路径