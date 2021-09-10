package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/router"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/env"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/logger"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/shutdown"

	"go.uber.org/zap"
)

// @host 127.0.0.1:8006
func main() {
	loggers, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.Get().LogPath()),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = loggers.Sync()
	}()

	s, err := router.NewHTTPServer(loggers)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggers.Fatal("http server startup err", zap.Error(err))
		}
	}()

	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				loggers.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					loggers.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					loggers.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					loggers.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 gRPC client
		func() {
			if s.GrpClient != nil {
				if err := s.GrpClient.Conn().Close(); err != nil {
					loggers.Error("gRPC client close err", zap.Error(err))
				}
			}
		},
	)
}
