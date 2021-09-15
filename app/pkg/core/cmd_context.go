package core

import (
	stdctx "context"
	"github.com/gin-gonic/gin"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errno"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/trace"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

type cmdContext struct {
	logger *zap.Logger
	trace Trace
}

func NewCmdContext(logger *zap.Logger) Context {
	return cmdContext{
		logger: logger,
	}
}

func (c cmdContext) init() {}

func (c cmdContext) ShouldBindQuery(obj interface{}) error {
	return nil
}

func (c cmdContext) ShouldBindPostForm(obj interface{}) error {
	return nil
}

func (c cmdContext) ShouldBindForm(obj interface{}) error {
	return nil
}

func (c cmdContext) ShouldBindJSON(obj interface{}) error {
	return nil
}

func (c cmdContext) ShouldBindURI(obj interface{}) error {
	return nil
}

func (c cmdContext) Redirect(code int, location string) {}

func (c cmdContext) Trace() Trace {
	if c.trace == nil {
		c.trace = trace.New("")
	}

	return c.trace
}

func (c cmdContext) setTrace(trace Trace) {
	c.trace = trace
}

func (c cmdContext) disableTrace() {}

func (c cmdContext) TraceId() string {
	if t := c.Trace(); t != nil {
		return t.ID()
	}
	return ""
}

func (c cmdContext) Logger() *zap.Logger {
	return c.logger
}

func (c cmdContext) setLogger(logger *zap.Logger) {
	c.logger = logger
}

func (c cmdContext) Info(msg string, fields ...zap.Field) {
	c.logger.Info(msg, fields...)
}

func (c cmdContext) Error(msg string, fields ...zap.Field) {
	c.logger.Error(msg, fields...)
}

func (c cmdContext) Warning(msg string, fields ...zap.Field) {
	c.logger.Warn(msg, fields...)
}

func (c cmdContext) Success(data interface{}) {}

func (c cmdContext) getSuccess() interface{} {
	return nil
}

func (c cmdContext) Failed(err errno.Error) {}

func (c cmdContext) failedError() errno.Error {
	return nil
}

func (c cmdContext) Header() http.Header {
	return nil
}

func (c cmdContext) GetHeader(key string) string {
	return ""
}

func (c cmdContext) SetHeader(key, value string) {}

func (c cmdContext) UserID() int64 {
	return 0
}

func (c cmdContext) setUserID(userID int64) {}

func (c cmdContext) UserName() string {
	return ""
}

func (c cmdContext) setUserName(userName string) {}

func (c cmdContext) Alias() string {
	return ""
}

func (c cmdContext) setAlias(path string) {}

func (c cmdContext) RequestInputParams() url.Values {
	return nil
}

func (c cmdContext) RequestPostFormParams() url.Values {
	return nil
}

func (c cmdContext) Request() *http.Request {
	return nil
}

func (c cmdContext) RawData() []byte {
	return []byte("")
}

func (c cmdContext) Method() string {
	return ""
}

func (c cmdContext) Host() string {
	return ""
}

func (c cmdContext) Path() string {
	return ""
}

func (c cmdContext) URI() string {
	return ""
}

func (c cmdContext) RequestContext() StdContext {
	return StdContext{
		stdctx.Background(),
		c.Trace(),
		c.Logger(),
	}
}

func (c cmdContext) ResponseWriter() gin.ResponseWriter {
	return nil
}

