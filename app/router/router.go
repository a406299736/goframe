package router

import (
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/pkg/grpc"
	"github.com/a406299736/goframe/app/pkg/metrics"
	"github.com/a406299736/goframe/app/router/middleware"
	"github.com/a406299736/goframe/pkg/errors"

	"go.uber.org/zap"
)

type resource struct {
	mux     core.Mux
	logger  *zap.Logger
	grpConn grpc.ClientConn
	middles middleware.Middleware
}

type Server struct {
	Mux       core.Mux
	GrpClient grpc.ClientConn
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

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
	r.middles = middleware.New()

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
	s.GrpClient = r.grpConn

	return s, nil
}
