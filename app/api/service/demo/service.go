package demo

import (
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/pkg/db"
	"github.com/a406299736/goframe/app/pkg/redis"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/repository/db-repo/demo"
)

var _ Service = (*service)(nil)

// TODO 可扩展接口
type Service interface {
	i()

	Detail(ctx core.Context) (info *demo.Test1, err errors.Er)
}

type service struct {
	db    db.Repo
	cache redis.Repo
}

type Service1 struct {
	db    db.Repo
	cache redis.Repo
}

// 方式1. 返回接口 推荐
func New(db db.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}

// 方式2. 返回结构体
func New1(db db.Repo, cache redis.Repo) *Service1 {
	return &Service1{
		db:    db,
		cache: cache,
	}
}
