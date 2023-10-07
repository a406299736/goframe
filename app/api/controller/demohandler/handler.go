package demohandler

import (
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/service/demo"
	"github.com/a406299736/goframe/configs"
	"github.com/a406299736/goframe/pkg/hash"
)

var _ Handler = (*handler)(nil)

// 扩展接口
type Handler interface {
	i()

	// Detail 个人信息
	// @Tags API.admin
	// @Router /api/admin/info [get]
	Detail() core.HandlerFunc
}

type handler struct {
	hashids     hash.Hash
	demoService demo.Service
}

type handler1 struct {
	hashids     hash.Hash
	demoService *demo.Service1
}

// 方式1. 返回接口 推荐
// cache 用不到时 传nil即可
func New() Handler {
	return &handler{
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		//hashids:     nil, // 用不到可忽略或赋值nil
		demoService: demo.New(),
	}
}

func (h *handler) i() {}

// 方式2. 返回结构体
func New1() *handler1 {
	return &handler1{
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		demoService: demo.New1(),
	}
}
