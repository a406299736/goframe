package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/code"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errno"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
)

func (m *middleware) Token(ctx core.Context) (userId int64, userName string, err errno.Error) {
	token, e := ctx.Cookie(configs.HeaderLoginToken)
	if e != nil {
		return 0, "", errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(e)
	}

	if token == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Token 参数"))

		return
	}

	split := strings.Split(token, ".")
	if len(split) < 2 {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("未获取登录信息"))
		return
	}

	userEncoding, er := base64.RawStdEncoding.DecodeString(split[1])
	if er != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("解析出错"))
		return
	}

	user := struct {
		Userid int64 `json:"userid"`
		Username string `json:"username"`
	}{}
	dt := json.NewDecoder(strings.NewReader(string(userEncoding)))
	dt.UseNumber()
	if e := dt.Decode(&user); e != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("解析出错"))
		return
	}

	if user.Userid <= 0 {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("请先登录"))

		return
	}

	userId = user.Userid
	userName = user.Username

	return
}
