package user_func

import (
	"github.com/a406299736/goframe/internal/repository/dbrepo/user"
	"github.com/a406299736/goframe/pkg/db"
)

func One(id int) (*User, error) {
	one, err := user.NewBuilder().WhereId("=", id).First(db.IDb.DbR)
	if err != nil {
		return nil, err
	}

	return New(one), nil
}

func All() ([]*User, error) {
	users, err := user.NewBuilder().QueryAll(db.IDb.DbR)
	if err != nil {
		return nil, err
	}

	return fmt(users), nil
}
