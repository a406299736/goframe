package demohandler

import (
	"fmt"
	"net/http"

	"github.com/a406299736/goframe/app/api/service/demo"
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/pkg/db"
	"github.com/a406299736/goframe/app/pkg/redis"
	"github.com/a406299736/goframe/pkg/errno"
	"github.com/a406299736/goframe/pkg/hash"
)

// 相当于 controller 层
type handler2 struct {
	cache   redis.Repo
	hashids hash.Hash
	db      db.Repo
}

type infoReq struct {
	Id int32 `form:"id"`
}

func New2(db db.Repo, cache redis.Repo) *handler2 {
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

		req := new(infoReq)
		t := c.ShouldBindForm(req)
		if t != nil {
			fmt.Printf("%v\n", t)
		}

		info, e := srv.Info(req.Id)
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
