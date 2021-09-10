package middleware

import "gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"

func (m *middleware) DisableLog() core.HandlerFunc {
	return func(c core.Context) {
		core.DisableTrace(c)
	}
}
