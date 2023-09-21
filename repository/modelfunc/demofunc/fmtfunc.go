package demofunc

import (
	"github.com/a406299736/goframe/app/pkg/code"
	"github.com/a406299736/goframe/pkg/errors"
	"github.com/a406299736/goframe/repository/dbrepo/demo"
)

// format

type DemoFunc struct {
	*demo.Test1
}

func New(x *demo.Test1) (*DemoFunc, errors.Er) {
	if x == nil {
		return nil, code.NilError
	}
	return &DemoFunc{x}, nil
}
