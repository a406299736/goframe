package demo

import (
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/repository/dbrepo/demo"
)

var _ Service = (*service)(nil)

// TODO 可扩展接口
type Service interface {
	i()

	Detail(ctx core.Context) (info *demo.Test1, err errors.Er)
}

type service struct {
}

type Service1 struct {
}

// 方式1. 返回接口 推荐
func New() Service {
	return &service{}
}

func (s *service) i() {}

// 方式2. 返回结构体
func New1() *Service1 {
	return &Service1{}
}
