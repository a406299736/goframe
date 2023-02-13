package middleware

import (
	"net/http"
	"time"

	"github.com/a406299736/goframe/app/pkg/code"
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/pkg/redis"
	"github.com/a406299736/goframe/configs"
	"github.com/a406299736/goframe/pkg/errno"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/pkg/token"
)

const reSubmitMark = "1"

func (m *middleware) Resubmit() core.HandlerFunc {
	return func(c core.Context) {
		cfg := configs.Get().URLToken

		tokenString, err := token.New(cfg.Secret).UrlSign(c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.Failed(errno.NewError(
				http.StatusBadRequest,
				code.UrlSignError,
				code.Text(code.UrlSignError)).WithErr(err),
			)
			return
		}

		redisKey := configs.RedisKeyPrefixRequestID + tokenString
		if !m.cache.Exists(redisKey) {
			err = m.cache.Set(redisKey, reSubmitMark, time.Minute*cfg.ExpireDuration)
			if err != nil {
				c.Failed(errno.NewError(
					http.StatusBadRequest,
					code.CacheSetError,
					code.Text(code.CacheSetError)).WithErr(err),
				)
				return
			}

			return
		}

		redisValue, err := m.cache.Get(redisKey, redis.WithTrace(c.Trace()))
		if err != nil {
			c.Failed(errno.NewError(
				http.StatusBadRequest,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(err),
			)
			return
		}

		if redisValue == reSubmitMark {
			c.Failed(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitError,
				code.Text(code.ResubmitError)).WithErr(errors.New("resubmit")),
			)
			return
		}

		return
	}
}
