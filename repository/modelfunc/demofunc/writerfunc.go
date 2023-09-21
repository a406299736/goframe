package demofunc

import (
	"github.com/a406299736/goframe/app/pkg/db"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/repository/dbrepo/demo"
)

// 写方法：insert，update，delete

func Insert(x *demo.Test1) (*DemoFunc, errors.Er) {
	id, err := x.Create(db.IDb.GetDbW())
	if err != nil {
		return nil, errors.NewErr(-1, err.Error())
	}

	x.Id = id
	return New(x)
}
