package demo_handler

import (
	"net/http"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/api/service/demo"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/db"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errno"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/hash"
)

// 相当于 controller 层
type handler2 struct {
	cache   cache.Repo
	hashids hash.Hash
	db      db.Repo
}

func New2(db db.Repo, cache cache.Repo) *handler2 {
	return &handler2{
		cache: cache,
		db:    db,
		//hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		//hashids:     nil, // 用不到可忽略或赋值nil
	}
}

func (h *handler2) Info() core.HandlerFunc {
	return func(c core.Context) {
		srv := demo.NewDemoService2(h.db, h.cache, c)
		info, e := srv.Info()
		if e != nil {
			c.Failed(errno.NewError(http.StatusOK, e.Code(), e.Error()))
		}
		c.Success(info)
	}
}

func (h *handler2) Create() core.HandlerFunc {
	return func(c core.Context) {
		srv := demo.NewDemoService2(h.db, h.cache, c)
		info, e := srv.Create()
		if e != nil {
			c.Failed(errno.NewError(http.StatusOK, e.Code(), e.Error()))
		}
		c.Success(info)
	}
}
