package user_func

import "github.com/a406299736/goframe/internal/repository/dbrepo/user"

type User struct {
	*user.User
}

func New(user *user.User) *User {
	if user == nil {
		return nil
	}

	return &User{user}
}

func fmt(users []*user.User) []*User {
	var data []*User
	for _, user := range users {
		if u := New(user); u != nil {
			data = append(data, u)
		}
	}

	return data
}
