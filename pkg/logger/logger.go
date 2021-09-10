package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/trace"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultLevel = zapcore.InfoLevel

	DefaultTimeLayout = time.RFC3339
)

type Option func(*option)

type option struct {
	level          zapcore.Level
	fields         map[string]string
	file           io.Writer
	timeLayout     string
	disableConsole bool
	isWithTrace    bool
}

func WithTrace() Option {
	return func(opt *option) {
		opt.isWithTrace = true
	}
}

// 输出debug日志
func WithDebugLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.DebugLevel
	}
}

// 输出Info日志
func WithInfoLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.InfoLevel
	}
}

// 输出warning日志
func WithWarnLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.WarnLevel
	}
}

// 输出error日志
func WithErrorLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.ErrorLevel
	}
}

func WithField(key, value string) Option {
	return func(opt *option) {
		opt.fields[key] = value
	}
}

// 输出到文件
func WithFileP(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}

	return func(opt *option) {
		opt.file = zapcore.Lock(f)
	}
}

func WithFileRotationP(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	return func(opt *option) {
		opt.file = &lumberjack.Logger{
			Filename:   file, // 文件路径
			MaxSize:    128,  // 单个文件最大尺寸，默认单位 M
			MaxBackups: 300,  // 最多保留 300 个备份
			MaxAge:     30,   // 最大时间，默认单位 day
			LocalTime:  true, // 使用本地时间
			Compress:   true, // 是否压缩 disabled by default
		}
	}
}

func WithTimeLayout(timeLayout string) Option {
	return func(opt *option) {
		opt.timeLayout = timeLayout
	}
}

func WithDisableConsole() Option {
	return func(opt *option) {
		opt.disableConsole = true
	}
}

func NewJSONLogger(opts ...Option) (*zap.Logger, error) {
	opt := &option{level: DefaultLevel, fields: make(map[string]string)}
	for _, f := range opts {
		f(opt)
	}

	timeLayout := DefaultTimeLayout
	if opt.timeLayout != "" {
		timeLayout = opt.timeLayout
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(timeLayout))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= opt.level && lvl < zapcore.ErrorLevel
	})

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= opt.level && lvl >= zapcore.ErrorLevel
	})

	stdout := zapcore.Lock(os.Stdout)
	stderr := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee()

	if !opt.disableConsole {
		core = zapcore.NewTee(
			zapcore.NewCore(jsonEncoder,
				zapcore.NewMultiWriteSyncer(stdout),
				lowPriority,
			),
			zapcore.NewCore(jsonEncoder,
				zapcore.NewMultiWriteSyncer(stderr),
				highPriority,
			),
		)
	}

	if opt.file != nil {
		core = zapcore.NewTee(core,
			zapcore.NewCore(jsonEncoder,
				zapcore.AddSync(opt.file),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= opt.level
				}),
			),
		)
	}

	logger := zap.New(core,
		zap.AddCaller(),
		zap.ErrorOutput(stderr),
	)

	for key, value := range opt.fields {
		logger = logger.WithOptions(zap.Fields(zapcore.Field{Key: key, Type: zapcore.StringType, String: value}))
	}

	// 多次New获取logger的 trace-id 会不一样
	// todo 解决上面id不一致的问题
	if opt.isWithTrace {
		tc := trace.New("")
		logger = logger.With(zap.Any("Trace-Id", tc.ID()))
	}

	return logger, nil
}

var _ Meta = (*meta)(nil)

type Meta interface {
	Key() string
	Value() interface{}
	meta()
}

type meta struct {
	key   string
	value interface{}
}

func (m *meta) Key() string {
	return m.key
}

func (m *meta) Value() interface{} {
	return m.value
}

func (m *meta) meta() {}

func NewMeta(key string, value interface{}) Meta {
	return &meta{key: key, value: value}
}

func WrapMeta(err error, metas ...Meta) (fields []zap.Field) {
	capacity := len(metas) + 1
	if err != nil {
		capacity++
	}

	fields = make([]zap.Field, 0, capacity)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	fields = append(fields, zap.Namespace("meta"))
	for _, meta := range metas {
		fields = append(fields, zap.Any(meta.Key(), meta.Value()))
	}

	return
}
