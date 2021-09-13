package core

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/code"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/env"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errno"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/trace"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cors "github.com/rs/cors/wrapper/gin"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

const _UI = `Go Basic`

const _MaxBurstSize = 100000

type Option func(*option)

type option struct {
	disablePProf      bool
	disablePrometheus bool
	panicNotify       OnPanicNotify
	recordMetrics     RecordMetrics
	enableCors        bool
	enableRate        bool
	enableOpenBrowser string
}

// OnPanicNotify 发生panic时通知用
type OnPanicNotify func(ctx Context, err interface{}, stackInfo string)

// RecordMetrics 记录prometheus指标用
// 如果使用AliasForRecordMetrics配置了别名，uri将被替换为别名。
type RecordMetrics func(method, uri string, success bool, httpCode, businessCode int, costSeconds float64, traceId string)

// WithDisablePProf 禁用 pprof 可视化和分析性能分析数据的工具
func WithDisablePProf() Option {
	return func(opt *option) {
		opt.disablePProf = true
	}
}

// WithDisablePrometheus 禁用prometheus 阿里云开源监控
func WithDisablePrometheus() Option {
	return func(opt *option) {
		opt.disablePrometheus = true
	}
}

// WithRecordMetrics 设置记录prometheus记录指标回调
func WithRecordMetrics(record RecordMetrics) Option {
	return func(opt *option) {
		opt.recordMetrics = record
	}
}

// WithEnableCors 开启CORS 跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// WithPanicNotify 设置panic时的通知消息
func WithPanicNotify(notify OnPanicNotify) Option {
	return func(opt *option) {
		opt.panicNotify = notify
	}
}

// WithEnableRate 是否限速
func WithEnableRate() Option {
	return func(opt *option) {
		opt.enableRate = true
	}
}

func DisableTrace(ctx Context) {
	ctx.disableTrace()
}

// AliasForRecordMetrics 对请求uri起个别名，用于prometheus记录指标。
// 如：Get /user/:username 这样的uri，因为username会有非常多的情况，这样记录prometheus指标会非常的不友好。
func AliasForRecordMetrics(path string) HandlerFunc {
	return func(ctx Context) {
		ctx.setAlias(path)
	}
}

// WrapAuthHandler 用来处理 Auth 的入口，在之后的handler中只需 ctx.UserID() ctx.UserName()
func WrapAuthHandler(handler func(Context) (userID int64, userName string, err errno.Error)) HandlerFunc {
	return func(ctx Context) {
		userID, userName, err := handler(ctx)
		if err != nil {
			ctx.Failed(err)
			return
		}
		ctx.setUserID(userID)
		ctx.setUserName(userName)
	}
}

// RouterGroup 包装gin的RouterGroup
type RouterGroup interface {
	Group(string, ...HandlerFunc) RouterGroup
	IRoutes
}

var _ IRoutes = (*router)(nil)

// IRoutes 封装gin IRoutes
type IRoutes interface {
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

type router struct {
	group *gin.RouterGroup
}

func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

func (r *router) Any(relativePath string, handlers ...HandlerFunc) {
	r.group.Any(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.group.DELETE(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	r.group.OPTIONS(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	r.group.HEAD(relativePath, wrapHandlers(handlers...)...)
}

func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)
			handler(ctx)
		}
	}

	return funcs
}

var _ Mux = (*mux)(nil)

type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func New(logger *zap.Logger, options ...Option) (Mux, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	gin.SetMode(gin.ReleaseMode)
	mux := &mux{
		engine: gin.New(),
	}

	// withoutLogPaths 不记录日志
	withoutTracePaths := map[string]bool{
		"/metrics": true,

		"/debug/pprof/":             true,
		"/debug/pprof/cmdline":      true,
		"/debug/pprof/profile":      true,
		"/debug/pprof/symbol":       true,
		"/debug/pprof/trace":        true,
		"/debug/pprof/allocs":       true,
		"/debug/pprof/block":        true,
		"/debug/pprof/goroutine":    true,
		"/debug/pprof/heap":         true,
		"/debug/pprof/mutex":        true,
		"/debug/pprof/threadcreate": true,

		"/favicon.ico": true,

		"/system/health": true,
	}

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	if !opt.disablePProf {
		if !env.Active().IsPro() {
			pprof.Register(mux.engine)
		}
	}

	if !opt.disablePrometheus {
		mux.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	if opt.enableCors {
		mux.engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("got panic one", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
		}()

		ctx.Next()
	})

	mux.engine.Use(func(ctx *gin.Context) {
		ts := time.Now()

		context := newContext(ctx)
		defer releaseContext(context)

		context.init()
		context.setLogger(logger)

		if !withoutTracePaths[ctx.Request.URL.Path] {
			if traceId := context.GetHeader(trace.Header); traceId != "" {
				context.setTrace(trace.New(traceId))
			} else {
				context.setTrace(trace.New(""))
			}
		}

		defer func() {
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())

				context.Error(
					"got panic two",
					zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo),
				)

				context.Failed(errno.NewError(
					http.StatusInternalServerError,
					code.ServerError,
					code.Text(code.ServerError)),
				)

				if notify := opt.panicNotify; notify != nil {
					notify(context, err, stackInfo)
				}
			}

			if ctx.Writer.Status() == http.StatusNotFound {
				return
			}

			var (
				response        interface{}
				businessCode    int
				businessCodeMsg string
				abortErr        error
				traceId         string
			)

			if ctx.IsAborted() {
				for i := range ctx.Errors { // 框架 error
					multierr.AppendInto(&abortErr, ctx.Errors[i])
				}

				if err := context.failedError(); err != nil { // 自定义 err
					multierr.AppendInto(&abortErr, err.GetErr())
					response = err
					businessCode = err.GetBusinessCode()
					businessCodeMsg = err.GetMsg()

					if x := context.Trace(); x != nil {
						context.SetHeader(trace.Header, x.ID())
						traceId = x.ID()
					}

					ctx.JSON(err.GetHttpCode(), &code.Failure{
						Code:    businessCode,
						Message: businessCodeMsg,
					})
				}
			} else {
				resp := context.getSuccess()
				response = struct {
					Code int         `json:"code"`
					Msg  string      `json:"msg"`
					Data interface{} `json:"data"`
				}{Code: businessCode, Msg: "success", Data: resp}

				if x := context.Trace(); x != nil {
					context.SetHeader(trace.Header, x.ID())
					traceId = x.ID()
				}

				ctx.JSON(http.StatusOK, response)
			}

			if opt.recordMetrics != nil {
				uri := context.URI()
				if alias := context.Alias(); alias != "" {
					uri = alias
				}

				opt.recordMetrics(
					context.Method(),
					uri,
					!ctx.IsAborted() && ctx.Writer.Status() == http.StatusOK,
					ctx.Writer.Status(),
					businessCode,
					time.Since(ts).Seconds(),
					traceId,
				)
			}

			var t *trace.Trace
			if x := context.Trace(); x != nil {
				t = x.(*trace.Trace)
			} else {
				return
			}

			decodedURL, _ := url.QueryUnescape(ctx.Request.URL.RequestURI())

			// ctx.Request.Header，精简 Header 参数
			traceHeader := map[string]string{
				"Content-Type":              ctx.GetHeader("Content-Type"),
				configs.HeaderLoginToken:    ctx.GetHeader(configs.HeaderLoginToken),
				configs.HeaderSignToken:     ctx.GetHeader(configs.HeaderSignToken),
				configs.HeaderSignTokenDate: ctx.GetHeader(configs.HeaderSignTokenDate),
			}

			t.WithRequest(&trace.Request{
				TTL:        "un-limit",
				Method:     ctx.Request.Method,
				DecodedURL: decodedURL,
				Header:     traceHeader,
				Body:       string(context.RawData()),
			})

			var responseBody interface{}

			if response != nil {
				responseBody = response
			}

			t.WithResponse(&trace.Response{
				Header:          ctx.Writer.Header(),
				HttpCode:        ctx.Writer.Status(),
				HttpCodeMsg:     http.StatusText(ctx.Writer.Status()),
				BusinessCode:    businessCode,
				BusinessCodeMsg: businessCodeMsg,
				Body:            responseBody,
				RTime:           time.Since(ts).Seconds(),
			})

			t.Success = !ctx.IsAborted() && ctx.Writer.Status() == http.StatusOK
			t.RTime = time.Since(ts).Seconds()

			logger.Info("core-interceptor",
				zap.Any("Trace-Id", t.Identifier),
				zap.Any("method", ctx.Request.Method),
				zap.Any("path", decodedURL),
				zap.Any("http_code", ctx.Writer.Status()),
				zap.Any("business_code", businessCode),
				zap.Any("success", t.Success),
				zap.Any("r_time", t.RTime),
				zap.Any("trace_info", t),
				zap.Error(abortErr),
			)
		}()

		ctx.Next()
	})

	if opt.enableRate {
		limiter := rate.NewLimiter(rate.Every(time.Second*1), _MaxBurstSize)
		mux.engine.Use(func(ctx *gin.Context) {
			context := newContext(ctx)
			defer releaseContext(context)

			if !limiter.Allow() {
				context.Failed(errno.NewError(
					http.StatusTooManyRequests,
					code.TooManyRequests,
					code.Text(code.TooManyRequests)),
				)
				return
			}

			ctx.Next()
		})
	}

	mux.engine.NoMethod(wrapHandlers(DisableTrace)...)
	mux.engine.NoRoute(wrapHandlers(DisableTrace)...)
	system := mux.Group("/system")
	{
		// 健康检查
		system.GET("/health", func(ctx Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Host(),
				Status:      "ok",
			}
			ctx.Success(resp)
		})
	}

	return mux, nil
}
