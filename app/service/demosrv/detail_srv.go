package demosrv

import (
	"github.com/a406299736/goframe/app/pkg/code"
	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/app/pkg/db"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/repository/dbrepo/demo"

	"go.uber.org/zap"
)

// 方式1
// demo 建议service不记录日志，错误直接返回上层，由app/pkg/core/core.go统一处理
func (src *service) Detail(c core.Context) (info *demo.Test1, err errors.Er) {
	info = &demo.Test1{}
	//err = errors.NewErr(code.UserNotExists, "msg error")

	i, _ := demo.NewDemoQueryBuilder().WhereIdIn([]int32{1}).First(db.IDb.GetDbR().WithContext(c.RequestContext()))
	c.Logger().Info("demo", zap.Any("demo", i.Id))

	dm := demo.NewDemoQueryBuilder()
	all, er := dm.WhereIdIn([]int32{1, 2, 3}).
		OrderById(false).
		QueryAll(db.IDb.GetDbR().WithContext(c.RequestContext()))
	if er != nil {
		err = errors.NewErr(code.UserNotExists, "query all error")
	}
	info = all[0]

	// 非必要情况也可以如下：
	//c.Info("demo", zap.Any("demo", "aaaaa1"))
	//c.Logger().Info("demo", zap.Any("demo", "bbbbbbb2"))

	return
}

// 方式2 不推荐
func (src *Service1) Detail1() (info *demo.Demo, err errors.Er) {
	info = &demo.Demo{}
	info.Username = "demo-1"
	info.Mobile = "123123131"
	info.Nickname = "nick name"

	return
}
