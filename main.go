package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/a406299736/goframe/configs"
	"github.com/a406299736/goframe/internal/router"
	db2 "github.com/a406299736/goframe/pkg/db"
	"github.com/a406299736/goframe/pkg/logger"
	"github.com/a406299736/goframe/pkg/redis"
	"github.com/a406299736/goframe/pkg/shutdown"
	"github.com/a406299736/goframe/pkg/tools"

	"go.uber.org/zap"
)

// @host 127.0.0.1:8006
func main() {
	loggers, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, configs.Get().App.Env)),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.Get().LogPath()),
	)
	if err != nil {
		panic(err)
	}

	// 关闭apollo
	//go func() {
	//	err = apollo.CheckStart()
	//	if err != nil {
	//		panic(err)
	//	}
	//}()

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

	tools.FmtPrintf("服务端口：%s", configs.ProjectPort)

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
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
			if db2.IDb != nil {
				if err := db2.IDb.DbWClose(); err != nil {
					loggers.Error("dbw close err", zap.Error(err))
				}

				if err := db2.IDb.DbRClose(); err != nil {
					loggers.Error("dbr close err", zap.Error(err))
				}
			}

			if db2.Conn1 != nil {
				if err := db2.Conn1.DbWClose(); err != nil {
					loggers.Error("dbw close err", zap.Error(err))
				}

				if err := db2.Conn1.DbRClose(); err != nil {
					loggers.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if redis.RedisRepo != nil {
				if err := redis.RedisRepo.Close(); err != nil {
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
