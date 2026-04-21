package middleware

import (
	core2 "github.com/a406299736/goframe/internal/core"
)

func (m *middleware) DisableLog() core2.HandlerFunc {
	return func(c core2.Context) {
		core2.DisableTrace(c)
	}
}
