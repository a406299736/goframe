package demofunc

import (
	"github.com/a406299736/goframe/app/pkg/code"
	"github.com/a406299736/goframe/app/pkg/db"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/repository/dbrepo/demo"
)

// 读方法：select

func OneBy(id int32) (*DemoFunc, errors.Er) {
	row, _ := demo.NewDemoQueryBuilder().First(db.IDb.GetDbR())
	return New(row)
}

func All() ([]*DemoFunc, errors.Er) {
	all, err := demo.NewDemoQueryBuilder().QueryAll(db.IDb.GetDbR())
	if err != nil {
		return nil, errors.NewErr(-1, err.Error())
	}

	return fmtRows(all)
}

func More() {

}

func fmtRows(rows []*demo.Test1) ([]*DemoFunc, errors.Er) {
	if rows == nil || len(rows) == 0 {
		return nil, code.NilError
	}

	f := make([]*DemoFunc, len(rows))
	for k := range rows {
		f[k], _ = New(rows[k])
	}
	return f, nil
}
