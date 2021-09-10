package middleware

import (
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errno"
)

func (m *middleware) Jwt(ctx core.Context) (userId int64, userName string, err errno.Error) {
	// TODO 暂未实现

	return
}
