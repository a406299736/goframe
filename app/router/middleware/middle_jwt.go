package middleware

import (
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/pkg/errno"
)

func (m *middleware) Jwt(ctx core.Context) (userId int64, userName string, err errno.Error) {
	// TODO 暂未实现

	return
}
