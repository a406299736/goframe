package errors

import (
	"bytes"
	"fmt"
)

type err struct {
	errType, errMsg, fullMsg string
	errCode                  int
	data                     []interface{}
}

type Er interface {
	Type() string
	Error() string
	Code() int
	Data() []interface{}
}

// service侧error
func NewErr(errCode int, errMsg string, dt ...interface{}) Er {
	return &err{errCode: errCode, errMsg: errMsg, data: dt}
}

func (e *err) Code() int {
	return e.errCode
}

func (e *err) Type() string {
	return e.errType
}

func (e *err) Error() string {
	if e.fullMsg == "" {
		e.genFullErrMsg()
	}
	return e.fullMsg
}

func (e *err) Data() []interface{} {
	if e.data == nil {
		e.data = make([]interface{}, 1)
	}
	return e.data
}

func (e *err) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Err: ")
	if e.errType != "" {
		buffer.WriteString(string(e.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(e.errMsg)
	e.fullMsg = fmt.Sprintf("%s\n", buffer.String())
	return
}
