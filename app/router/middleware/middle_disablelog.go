package middleware

import "github.com/a406299736/goframe/app/pkg/core"

func (m *middleware) DisableLog() core.HandlerFunc {
	return func(c core.Context) {
		core.DisableTrace(c)
	}
}
