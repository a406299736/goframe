package middleware

import (
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
)

func (m *middleware) Signature() core.HandlerFunc {
	return func(c core.Context) {
		// TODO 暂未实现

		return
	}
}
