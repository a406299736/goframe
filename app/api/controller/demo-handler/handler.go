package demo_handler

import (
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/api/service/demo"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/db"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/hash"
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
	cache       cache.Repo
	hashids     hash.Hash
	demoService demo.Service
}

type handler1 struct {
	cache       cache.Repo
	hashids     hash.Hash
	demoService *demo.Service1
}

// 方式1. 返回接口 推荐
// cache 用不到时 传nil即可
func New(db db.Repo, cache cache.Repo) Handler {
	return &handler{
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		//hashids:     nil, // 用不到可忽略或赋值nil
		demoService: demo.New(db, cache),
	}
}

func (h *handler) i() {}

// 方式2. 返回结构体
func New1(db db.Repo, cache cache.Repo) *handler1 {
	return &handler1{
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		demoService: demo.New1(db, cache),
	}
}
