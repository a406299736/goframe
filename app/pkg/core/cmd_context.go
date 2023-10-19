package core

import (
	stdctx "context"
	"net/http"
	"net/url"

	"github.com/a406299736/goframe/pkg/errno"
	"github.com/a406299736/goframe/pkg/trace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Context = (*cmdContext)(nil)

type cmdContext struct {
	Context
	logger *zap.Logger
	trace  Trace
}

func NewCmdContext(logger *zap.Logger) Context {
	return &cmdContext{
		logger: logger,
	}
}

func (c *cmdContext) init() {}

func (c *cmdContext) ShouldBindQuery(obj interface{}) error {
	return nil
}

func (c *cmdContext) ShouldBindPostForm(obj interface{}) error {
	return nil
}

func (c *cmdContext) ShouldBindForm(obj interface{}) error {
	return nil
}

func (c *cmdContext) ShouldBindJSON(obj interface{}) error {
	return nil
}

func (c *cmdContext) ShouldBindURI(obj interface{}) error {
	return nil
}

func (c *cmdContext) ShouldBind(obj interface{}) error {
	return nil
}

func (c *cmdContext) Redirect(code int, location string) {}

func (c *cmdContext) Trace() Trace {
	if c.trace == nil {
		c.trace = trace.New("")
	}

	return c.trace
}

func (c *cmdContext) setTrace(trace Trace) {
	c.trace = trace
}

func (c *cmdContext) disableTrace() {}

func (c *cmdContext) TraceId() string {
	if t := c.Trace(); t != nil {
		return t.ID()
	}
	return ""
}

func (c *cmdContext) Logger() *zap.Logger {
	return c.logger
}

func (c *cmdContext) setLogger(logger *zap.Logger) {
	c.logger = logger
}

func (c *cmdContext) AnyInfo(msg, key string, val any) {
	c.Logger().Info(msg, zap.Any(key, val))
}

func (c *cmdContext) AnyError(msg, key string, val any) {
	c.Logger().Error(msg, zap.Any(key, val))
}

func (c *cmdContext) AnyWarning(msg, key string, val any) {
	c.Logger().Warn(msg, zap.Any(key, val))
}
func (c *cmdContext) Success(data interface{}) {}

func (c *cmdContext) getSuccess() interface{} {
	return nil
}

func (c *cmdContext) Failed(err errno.Error) {}

func (c *cmdContext) failedError() errno.Error {
	return nil
}

func (c *cmdContext) Header() http.Header {
	return nil
}

func (c *cmdContext) GetHeader(key string) string {
	return ""
}

func (c *cmdContext) SetHeader(key, value string) {}

func (c *cmdContext) UserID() int64 {
	return 0
}

func (c *cmdContext) setUserID(userID int64) {}

func (c *cmdContext) UserName() string {
	return ""
}

func (c *cmdContext) setUserName(userName string) {}

func (c *cmdContext) Alias() string {
	return ""
}

func (c *cmdContext) setAlias(path string) {}

func (c *cmdContext) RequestInputParams() url.Values {
	return nil
}

func (c *cmdContext) RequestPostFormParams() url.Values {
	return nil
}

func (c *cmdContext) Request() *http.Request {
	return nil
}

func (c *cmdContext) RawData() []byte {
	return []byte("")
}

func (c *cmdContext) Method() string {
	return ""
}

func (c *cmdContext) Host() string {
	return ""
}

func (c *cmdContext) Path() string {
	return ""
}

func (c *cmdContext) URI() string {
	return ""
}

func (c *cmdContext) RequestContext() StdContext {
	return StdContext{
		stdctx.Background(),
		c.Trace(),
		c.Logger(),
	}
}

func (c *cmdContext) ResponseWriter() gin.ResponseWriter {
	return nil
}
