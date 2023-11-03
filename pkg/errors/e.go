package errors

type err struct {
	b       bool
	errCode int
	errMsg  string
	data    []any
}

type option func(*err)

type Er interface {
	Empty() bool
	Error() string
	Code() int
	Data() []interface{}
}

func WithData(data ...any) option {
	return func(e *err) {
		e.data = data
	}
}

func WithEmpty(b bool) option {
	return func(e *err) {
		e.b = b
	}
}

// service侧error
func NewErr(errCode int, errMsg string, opts ...option) Er {
	e := &err{errCode: errCode, errMsg: errMsg}
	for _, opt := range opts {
		opt(e)
	}

	return e
}

func (e *err) Code() int {
	return e.errCode
}

func (e *err) Empty() bool {
	return e.b
}

func (e *err) Error() string {
	return e.errMsg
}

func (e *err) Data() []interface{} {
	if e.data == nil {
		e.data = make([]interface{}, 1)
	}
	return e.data
}
