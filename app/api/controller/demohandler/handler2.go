package demohandler

import (
	"fmt"
	"net/http"

	"github.com/a406299736/goframe/app/api/service/demo"
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/pkg/errno"
	"github.com/a406299736/goframe/pkg/hash"
)

// 相当于 controller 层
type handler2 struct {
	hashids hash.Hash
}

type infoReq struct {
	Id int32 `form:"id"`
}

func New2() *handler2 {
	return &handler2{
		//hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		//hashids:     nil, // 用不到可忽略或赋值nil
	}
}

func (h *handler2) Info() core.HandlerFunc {
	return func(c core.Context) {
		srv := demo.NewDemoService2()

		req := new(infoReq)
		t := c.ShouldBindForm(req)
		if t != nil {
			fmt.Printf("%v\n", t)
		}

		info, e := srv.Info(c, req.Id)
		if e != nil {
			c.Failed(errno.NewError(http.StatusOK, e.Code(), e.Error()))
		}
		c.Success(info)
	}
}

func (h *handler2) Create() core.HandlerFunc {
	return func(c core.Context) {
		srv := demo.NewDemoService2()
		info, e := srv.Create(c)
		if e != nil {
			c.Failed(errno.NewError(http.StatusOK, e.Code(), e.Error()))
		}
		c.Success(info)
	}
}
