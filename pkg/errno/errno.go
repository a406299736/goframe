package errno

import (
	"encoding/json"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
)

var _ Error = (*err)(nil)

type Error interface {
	// i 为了避免被其他包实现
	i()
	WithErr(err error) Error
	GetBusinessCode() int
	GetHttpCode() int
	GetMsg() string
	GetErr() error
	ToString() string
}

type err struct {
	HttpCode     int    // HTTP Code
	BusinessCode int    // 业务 Code
	Message      string // 描述信息
	Err          error  // 错误信息
}

func NewError(httpCode, businessCode int, msg string) Error {
	return &err{
		HttpCode:     httpCode,
		BusinessCode: businessCode,
		Message:      msg,
	}
}

func (e *err) i() {}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMsg() string {
	return e.Message
}

func (e *err) GetErr() error {
	return e.Err
}

func (e *err) ToString() string {
	err := &struct {
		HttpCode     int    `json:"http_code"`
		BusinessCode int    `json:"business_code"`
		Message      string `json:"message"`
	}{
		HttpCode:     e.HttpCode,
		BusinessCode: e.BusinessCode,
		Message:      e.Message,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
