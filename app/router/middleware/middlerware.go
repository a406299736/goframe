package middleware

import (
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/pkg/errno"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	// i 为了避免被其他包实现
	i()

	// Jwt 中间件
	Jwt(ctx core.Context) (userId int64, userName string, err errno.Error)

	// Resubmit 中间件
	Resubmit() core.HandlerFunc

	// DisableLog 不记录日志
	DisableLog() core.HandlerFunc

	// Signature 签名验证，对用签名算法 pkg/signature
	Signature() core.HandlerFunc

	// Token 签名验证，对登录用户的验证
	Token(ctx core.Context) (userId int64, userName string, err errno.Error)

	// RBAC 权限验证
	RBAC() core.HandlerFunc
}

type middleware struct {
}

func New() Middleware {
	return &middleware{}
}

func (m *middleware) i() {}
