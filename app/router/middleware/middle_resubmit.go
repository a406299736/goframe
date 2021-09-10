package middleware

import (
	"net/http"
	"time"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/code"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errno"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/token"
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

		redisValue, err := m.cache.Get(redisKey, cache.WithTrace(c.Trace()))
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
