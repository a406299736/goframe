package demo

import (
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/db"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/repository/db-repo/demo"
)

var _ Service = (*service)(nil)

// TODO 可扩展接口
type Service interface {
	i()

	Detail(ctx core.Context) (info *demo.WmAbout, err errors.Er)
}

type service struct {
	db    db.Repo
	cache cache.Repo
}

type Service1 struct {
	db    db.Repo
	cache cache.Repo
}

// 方式1. 返回接口 推荐
func New(db db.Repo, cache cache.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}

// 方式2. 返回结构体
func New1(db db.Repo, cache cache.Repo) *Service1 {
	return &Service1{
		db:    db,
		cache: cache,
	}
}
