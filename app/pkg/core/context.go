package core

import (
	"bytes"
	stdctx "context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/a406299736/goframe/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type HandlerFunc func(c Context)

const (
	_Alias          = "_alias_"
	_TraceName      = "_trace_"
	_LoggerName     = "_logger_"
	_BodyName       = "_body_"
	_SuccessName    = "_success _"
	_UserID         = "_user_id_"
	_UserName       = "_user_name_"
	_AbortErrorName = "_abort_error_"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(context)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*context)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*context)
	c.ctx = nil
	contextPool.Put(c)
}

var _ Context = (*context)(nil)

type Context interface {
	init()

	Trace() Trace
	setTrace(trace Trace)
	disableTrace()
	TraceId() string

	Logger() *zap.Logger
	setLogger(logger *zap.Logger)
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Warning(msg string, fields ...zap.Field)

	// Success 成功返回
	Success(data interface{})
	getSuccess() interface{}

	// Failed 错误返回
	Failed(err errno.Error)
	failedError() errno.Error

	Header() http.Header
	GetHeader(key string) string
	SetHeader(key, value string)

	Cookie(key string) (string, error)
	SetCookie(key, value string, maxAge int, path, domain string, secure, httpOnly bool)

	UserID() int64
	setUserID(userID int64)

	UserName() string
	setUserName(userName string)

	Alias() string
	setAlias(path string)

	// 反序列化 querystring
	// tag: `form:"xxx"` (注：不要写成 query)
	ShouldBindQuery(obj interface{}) error
	// 反序列化 postform (querystring会被忽略)
	// 兼容content-type类型：multipart/form-data, application/x-www-form-urlencoded, application/json;
	// tag: `json:"xxx" form:"xxx"`
	ShouldBindPostForm(obj interface{}) error
	// 同时反序列化 querystring 和 postform;
	// 当 querystring 和 postform 存在相同字段时，postform 优先使用。
	// tag: `form:"xxx"`
	ShouldBindForm(obj interface{}) error
	// 反序列化 postjson
	// tag: `json:"xxx"`
	ShouldBindJSON(obj interface{}) error
	// 反序列化 path 参数(如路由路径为 /user/:name)
	// tag: `uri:"xxx"`
	ShouldBindURI(obj interface{}) error
	ShouldBind(obj interface{}) error
	// 重定向
	Redirect(code int, location string)

	RequestInputParams() url.Values
	RequestPostFormParams() url.Values
	Request() *http.Request
	RawData() []byte
	Method() string
	Host() string
	Path() string
	URI() string
	RequestContext() StdContext

	ResponseWriter() gin.ResponseWriter
}

type context struct {
	ctx *gin.Context
}

type StdContext struct {
	stdctx.Context
	Trace
	*zap.Logger
}

func (c *context) init() {
	body, err := c.ctx.GetRawData()
	if err != nil {
		panic(err)
	}

	c.ctx.Set(_BodyName, body)
	c.ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
}

func (c *context) Cookie(key string) (string, error) {
	return c.ctx.Cookie(key)
}

func (c *context) SetCookie(key, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	c.ctx.SetCookie(key, value, maxAge, path, domain, secure, httpOnly)
}

func (c *context) ShouldBindQuery(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Query)
}

func (c *context) ShouldBindPostForm(obj interface{}) error {
	switch c.ctx.ContentType() {
	case binding.MIMEJSON:
		return c.ctx.ShouldBindWith(obj, binding.JSON)
	case binding.MIMEPOSTForm:
		return c.ctx.ShouldBindWith(obj, binding.Form)
	default:
		return c.ctx.ShouldBindWith(obj, binding.FormMultipart)
	}
}

func (c *context) ShouldBindForm(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.Form)
}

func (c *context) ShouldBindJSON(obj interface{}) error {
	return c.ctx.ShouldBindWith(obj, binding.JSON)
}

func (c *context) ShouldBindURI(obj interface{}) error {
	return c.ctx.ShouldBindUri(obj)
}

func (c *context) ShouldBind(obj interface{}) error {
	return c.ctx.ShouldBind(obj)
}

func (c *context) Redirect(code int, location string) {
	c.ctx.Redirect(code, location)
}

func (c *context) Trace() Trace {
	t, ok := c.ctx.Get(_TraceName)
	if !ok || t == nil {
		return nil
	}

	return t.(Trace)
}

func (c *context) setTrace(trace Trace) {
	c.ctx.Set(_TraceName, trace)
}

func (c *context) disableTrace() {
	c.setTrace(nil)
}

func (c *context) Logger() *zap.Logger {
	logger, ok := c.ctx.Get(_LoggerName)
	if !ok {
		return nil
	}

	return logger.(*zap.Logger).With(zap.Any("Trace-Id", c.TraceId()))
}

func (c *context) Info(msg string, fields ...zap.Field) {
	c.Logger().Info(msg, fields...)
}

func (c *context) Error(msg string, fields ...zap.Field) {
	c.Logger().Error(msg, fields...)
}

func (c *context) Warning(msg string, fields ...zap.Field) {
	c.Logger().Warn(msg, fields...)
}

func (c *context) setLogger(logger *zap.Logger) {
	c.ctx.Set(_LoggerName, logger)
}

func (c *context) getSuccess() interface{} {
	if success, ok := c.ctx.Get(_SuccessName); ok != false {
		return success
	}
	return nil
}

func (c *context) Success(success interface{}) {
	c.ctx.Set(_SuccessName, success)
}

func (c *context) Header() http.Header {
	header := c.ctx.Request.Header

	clone := make(http.Header, len(header))
	for k, v := range header {
		value := make([]string, len(v))
		copy(value, v)

		clone[k] = value
	}
	return clone
}

func (c *context) GetHeader(key string) string {
	return c.ctx.GetHeader(key)
}

func (c *context) SetHeader(key, value string) {
	c.ctx.Header(key, value)
}

func (c *context) UserID() int64 {
	val, ok := c.ctx.Get(_UserID)
	if !ok {
		return 0
	}

	return val.(int64)
}

func (c *context) setUserID(userID int64) {
	c.ctx.Set(_UserID, userID)
}

func (c *context) UserName() string {
	val, ok := c.ctx.Get(_UserName)
	if !ok {
		return ""
	}

	return val.(string)
}

func (c *context) setUserName(userName string) {
	c.ctx.Set(_UserName, userName)
}

func (c *context) Failed(err errno.Error) {
	if err != nil {
		httpCode := err.GetHttpCode()
		if httpCode == 0 {
			httpCode = http.StatusInternalServerError
		}

		c.ctx.AbortWithStatus(httpCode)
		c.ctx.Set(_AbortErrorName, err)
	}
}

func (c *context) failedError() errno.Error {
	err, _ := c.ctx.Get(_AbortErrorName)
	return err.(errno.Error)
}

func (c *context) Alias() string {
	path, ok := c.ctx.Get(_Alias)
	if !ok {
		return ""
	}

	return path.(string)
}

func (c *context) setAlias(path string) {
	if path = strings.TrimSpace(path); path != "" {
		c.ctx.Set(_Alias, path)
	}
}

func (c *context) RequestInputParams() url.Values {
	_ = c.ctx.Request.ParseForm()
	return c.ctx.Request.Form
}

func (c *context) RequestPostFormParams() url.Values {
	_ = c.ctx.Request.ParseForm()
	return c.ctx.Request.PostForm
}

func (c *context) Request() *http.Request {
	return c.ctx.Request
}

func (c *context) RawData() []byte {
	body, ok := c.ctx.Get(_BodyName)
	if !ok {
		return nil
	}

	return body.([]byte)
}

func (c *context) Method() string {
	return c.ctx.Request.Method
}

func (c *context) Host() string {
	return c.ctx.Request.Host
}

func (c *context) Path() string {
	return c.ctx.Request.URL.Path
}

func (c *context) URI() string {
	uri, _ := url.QueryUnescape(c.ctx.Request.URL.RequestURI())
	return uri
}

func (c *context) RequestContext() StdContext {
	return StdContext{
		stdctx.Background(),
		c.Trace(),
		c.Logger(),
	}
}

func (c *context) ResponseWriter() gin.ResponseWriter {
	return c.ctx.Writer
}

func (c *context) TraceId() string {
	if t := c.Trace(); t != nil {
		return t.ID()
	}
	return ""
}
