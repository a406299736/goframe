package demo

import (
	"encoding/json"
	"fmt"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/code"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/db"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/httpclient"
	db_repo "gitlab.weimiaocaishang.com/weimiao/go-basic/repository/db-repo"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/repository/db-repo/demo"
	"time"

	"go.uber.org/zap"
)

// 相当于 service 层
type service2 struct {
	db    db.Repo
	cache cache.Repo
	ctx   core.Context
}

func NewDemoService2(db db.Repo, cache cache.Repo, ctx core.Context) *service2 {
	return &service2{
		db:    db,
		cache: cache,
		ctx:   ctx,
	}
}

// 新增 省略传参...
func (s *service2) Create() (id int, e errors.Er) {
	demo2 := demo.NewDemo()
	demo2.Call = "call"
	demo2.Aspirations = "asp"
	demo2.Recruit = "rec"
	demo2.Map = "map"
	id, err := demo2.Create(s.db.GetDbR().WithContext(s.ctx.RequestContext()))
	if err != nil {
		return id, errors.NewErr(code.MySQLExecError, err.Error())
	}

	return id, nil
}

// 查询 省略传参...
func (s *service2) Info() (one *demo.WmAbout, e errors.Er) {
	s.ctx.Logger().Info("info", zap.Any("aaa", "bbbb"))

	demo2 := demo.NewDemoQueryBuilder()

	err2 := s.cache.Set("aaaa", "abs", time.Minute * 3, cache.WithTrace(s.ctx.Trace()))
	if err2 != nil {
		return nil, errors.NewErr(100000, err2.Error())
	}

	js, _ := json.Marshal(map[string]interface{}{"id":11001694})
	res, err := httpclient.PostJSON(configs.Get().Center.ClassUrl + "/api/class/info",
		js, httpclient.WithHeader("h-app-id", "98"))
	if err != nil {
		return nil, errors.NewErr(100002, err.Error())
	}
	fmt.Printf("%v", string(res))

	body, err2 := httpclient.Get("http://www.baidu.com", nil)
	if err2 != nil {
		return nil, errors.NewErr(100003, err2.Error())
	}
	fmt.Printf("%v", string(body))

	one, err = demo2.WhereId(db_repo.EqualPredicate, 1).
		QueryOne(s.db.GetDbR().WithContext(s.ctx.RequestContext()))
	if err != nil {
		return nil, errors.NewErr(code.MySQLExecError, err.Error())
	}

	return one, nil
}

// 更新 省略传参...
func (s *service2) Update2() errors.Er {
	demo2 := demo.NewDemoQueryBuilder()
	err := demo2.WhereMobile(db_repo.EqualPredicate, "18686868686").
		WhereNicknameIn([]string{"zhang.san", "li.si"}).
		WhereIsDeletedNotIn([]int32{2, 3, 4}).
		Updates(s.db.GetDbR().WithContext(s.ctx.RequestContext()),
			map[string]interface{}{"mobile": "1861021234", "nickname": "hahahha"})
	if err != nil {
		return errors.NewErr(code.MySQLExecError, err.Error())
	}

	return nil
}

// 删除 省略传参...
func (s *service2) Del() errors.Er {
	demo2 := demo.NewDemoQueryBuilder()
	err := demo2.WhereIdIn([]int32{1000, 2000, 3000}).Delete(s.db.GetDbR().WithContext(s.ctx.RequestContext()))
	if err != nil {
		return errors.NewErr(code.MySQLExecError, err.Error())
	}

	return nil
}

// 列表
func (s *service2) List() (lt []*demo.WmAbout, er errors.Er) {
	demo2 := demo.NewDemoQueryBuilder()
	all, err := demo2.WhereIsDeleted(db_repo.EqualPredicate, 0).
		OrderByUpdatedAt(false).
		Limit(1).Offset(200).
		QueryAll(s.db.GetDbR().WithContext(s.ctx.RequestContext()))
	if err != nil {
		return lt, errors.NewErr(code.MySQLExecError, err.Error())
	}

	return all, nil
}
