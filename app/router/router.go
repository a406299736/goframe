package router

import (
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/db"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/grpc"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/metrics"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/router/middleware"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"

	"go.uber.org/zap"
)

type resource struct {
	mux     core.Mux
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	grpConn grpc.ClientConn
	middles middleware.Middleware
}

type Server struct {
	Mux       core.Mux
	Db        db.Repo
	Cache     cache.Repo
	GrpClient grpc.ClientConn
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

	dbRepo, err := db.New()
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	r.db = dbRepo

	cacheRepo, err := cache.New()
	if err != nil {
		logger.Fatal("new cache err", zap.Error(err))
	}
	r.cache = cacheRepo

	gRPCRepo, err := grpc.New()
	if err != nil {
		logger.Fatal("new grpc err", zap.Error(err))
	}
	r.grpConn = gRPCRepo

	mux, err := core.New(logger,
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithRecordMetrics(metrics.RecordMetrics),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	r.middles = middleware.New(r.cache, r.db)

	// 注册 分组路由
	{
		// demo
		setDemoRouter(r)
		// demo2
		setDemo2Router2(r)
		// demo3 不传任何参数
		//setDemoRouter3() r.db 和 r.cache 都可以不传
	}

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache
	s.GrpClient = r.grpConn

	return s, nil
}
