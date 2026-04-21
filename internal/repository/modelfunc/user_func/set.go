package user_func

import (
	"github.com/a406299736/goframe/internal/repository/dbrepo/user"
	"github.com/a406299736/goframe/pkg/db"
)

func UpdateById(id int, data map[string]any) error {
	return user.NewBuilder().WhereId("=", id).Updates(db.IDb.DbW, data)
}

func Create(u *user.User) (int, error) {
	return u.Create(db.IDb.DbW)
}
