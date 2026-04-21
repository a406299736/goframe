package user_handler

import (
	"net/http"

	"github.com/a406299736/goframe/internal/core"
	"github.com/a406299736/goframe/internal/service/user_svc"
	"github.com/a406299736/goframe/pkg/errno"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (*UserHandler) Info() core.HandlerFunc {
	return func(c core.Context) {
		info, err := user_svc.Info(int(c.UserID()))
		if err != nil {
			c.Failed(errno.NewError(http.StatusOK, 1, err.Error()))
			return
		}

		c.Success(info)
	}
}
