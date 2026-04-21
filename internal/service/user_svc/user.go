package user_svc

import "github.com/a406299736/goframe/internal/repository/modelfunc/user_func"

func Info(id int) (info map[string]any, err error) {
	one, err := user_func.One(id)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":   one.Id,
		"name": one.Name,
	}, nil
}
