package demo

import (
	"encoding/json"
	"fmt"
	"github.com/a406299736/goframe/app/api/center"
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/pkg/db"
	"github.com/a406299736/goframe/pkg/apollo"
	"github.com/a406299736/goframe/pkg/errors"
	db_repo "github.com/a406299736/goframe/repository/dbrepo"
	"github.com/a406299736/goframe/repository/dbrepo/test1"
	"go.uber.org/zap"
	"strconv"
)

// 相当于 service 层
type service2 struct {
}

// 根据实际业务 参数自定义
func NewDemoService2() *service2 {
	return &service2{}
}

// 新增 省略传参...
func (s *service2) Create(ctx core.Context) (id int32, e errors.Er) {
	demo2 := test1.NewModel()
	demo2.Call = "call"
	demo2.Aspirations = "asp"
	demo2.Recruit = "rec"
	demo2.Map = "map"
	id, e = demo2.Create(db.IDb.GetDbR().WithContext(ctx.RequestContext()))
	if e != nil {
		return 0, e
	}

	return id, nil
}

// 查询 省略传参...
func (s *service2) Info(ctx core.Context, id int32) (one *test1.Test1, e errors.Er) {
	ctx.Logger().Info("info", zap.Any("aaa", "bbbb"))
	ctx.Logger().Info("user id:" + strconv.Itoa(int(ctx.UserID())))

	// apollo demo
	conf, err := apollo.New(apollo.WithNamespace("application"))
	if err != nil {
		fmt.Printf("%+v", err)
	}
	ctx.Info("apollo value USER_LIST:" + conf.GetStringValue("USER_LIST", "2222"))

	//fmt.Printf("%p\n", s.ctx)

	demo2 := test1.NewQueryBuilder()

	// redis demo
	//err2 := s.cache.Set("aaaa", "abs", time.Minute*3, cache.WithTrace(s.ctx.Trace()))
	//if err2 != nil {
	//	return nil, errors.NewErr(100000, err2.Error())
	//}

	//  http post请求中心demo
	js, _ := json.Marshal(map[string]interface{}{"id": 11001694})
	_, error2 := center.ClassInfo(js)
	if error2 != nil {
		return nil, error2
	}
	//fmt.Printf("%v", info)

	// http get demo
	//body, err2 := httpclient.Get("http://www.baidu.com", nil)
	//if err2 != nil {
	//	return nil, errors.NewErr(100003, err2.Error())
	//}
	//fmt.Printf("%v", string(body))

	// 查询db demo
	one, e = demo2.WhereId(db_repo.EPd, 1).
		QueryOne(db.IDb.GetDbR().WithContext(ctx.RequestContext()))
	if e != nil {
		return nil, e
	}

	return one, nil
}

// 更新 省略传参...
func (s *service2) Update2(ctx core.Context) errors.Er {
	demo2 := test1.NewQueryBuilder()
	err := demo2.WhereMap(db_repo.EPd, "18686868686").
		WhereRecruitIn([]string{"zhang.san", "li.si"}).
		Updates(db.IDb.GetDbR().WithContext(ctx.RequestContext()),
			map[string]interface{}{"aspirations": "aaaaaassss", "call": "hahahha"})
	if err != nil {
		return err
	}

	return nil
}

// 删除 省略传参...
func (s *service2) Del(ctx core.Context) errors.Er {
	demo2 := test1.NewQueryBuilder()
	err := demo2.WhereIdIn([]int32{1000, 2000, 3000}).Delete(db.IDb.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	return nil
}

// 列表
func (s *service2) List(ctx core.Context) (lt []*test1.Test1, er errors.Er) {
	demo2 := test1.NewQueryBuilder()
	lt, er = demo2.WhereAspirations(db_repo.EPd, "aaaa").
		OrderByUpdated(false).
		Limit(1).Offset(200).
		QueryAll(db.IDb.GetDbR().WithContext(ctx.RequestContext()))

	return
}
