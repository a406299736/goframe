package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a406299736/goframe/configs"
	"github.com/a406299736/goframe/internal/core"
	"github.com/a406299736/goframe/pkg/errno"
	"github.com/a406299736/goframe/pkg/redis"
	"github.com/a406299736/goframe/pkg/token"
)

const reSubmitMark = "1"

func (m *middleware) Resubmit() core.HandlerFunc {
	return func(c core.Context) {
		cfg := configs.Get().URLToken

		cache := redis.RedisRepo

		tokenString, err := token.New(cfg.Secret).UrlSign(c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.Failed(errno.NewError(
				http.StatusBadRequest,
				errno.SignError,
				errno.Text(errno.SignError)).WithErr(err),
			)
			return
		}

		redisKey := configs.RedisKeyPrefixRequestID + tokenString
		if !cache.Exists(redisKey) {
			err = cache.Set(redisKey, reSubmitMark, time.Minute*cfg.ExpireDuration)
			if err != nil {
				c.Failed(errno.NewError(
					http.StatusBadRequest,
					errno.CacheSetError,
					errno.Text(errno.CacheSetError)).WithErr(err),
				)
				return
			}

			return
		}

		redisValue, err := cache.Get(redisKey, redis.WithTrace(c.Trace()))
		if err != nil {
			c.Failed(errno.NewError(
				http.StatusBadRequest,
				errno.CacheGetError,
				errno.Text(errno.CacheGetError)).WithErr(err),
			)
			return
		}

		if redisValue == reSubmitMark {
			c.Failed(errno.NewError(
				http.StatusBadRequest,
				errno.ResubmitError,
				errno.Text(errno.ResubmitError)).WithErr(fmt.Errorf("resubmit")),
			)
			return
		}

		return
	}
}
