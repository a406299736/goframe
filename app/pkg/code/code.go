package code

import (
	"github.com/a406299736/goframe/configs"
	"github.com/a406299736/goframe/pkg/errno"
	"net/http"
)

type Failure struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	Success            = 0     // 成功code
	ServerError        = 11000 // 内部服务错误
	TooManyRequests    = 11001 // 请求频繁
	ResubmitError      = 11002 // 重复提交
	NotAllowed         = 12002 // 无权限
	AuthorizationError = 12003 // 授权错误
	SignError          = 12004 // 参数签名错误
	EncodeError        = 14006 // 加密失败
	DecodeError        = 14007 // 解密失败

	NotExists     = 20001 // 不存在
	KeyNotExists  = 20002 // key不存在
	UserNotExists = 20003 // 用户不存在
	ParseError    = 21000 // 解析失败
	ParamError    = 21001 // 参数错误
	ValueIsNil    = 21002 // nil错误

	ThirdRespError = 30000 // 三方服务返回错误

	QueryNotExist     = 50000 // 查询记录不存在
	MySQLConnectError = 50001 // 数据库连接错误
	MySQLExecError    = 50002 // sql执行错误
	UniqueKeyConflict = 50003 // 唯一键冲突
	CacheSetError     = 51000 // 缓存设置出错
	CacheGetError     = 51001 // 获取缓存出错
	CacheDelError     = 51002 // 删除缓存出错
	CacheNotExist     = 51003 // 缓存不存在
	RedisConnectError = 51004 // redis连接错误
)

func Text(code int) string {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		return zhCNText[code]
	}

	return zhCNText[code]
}

func RespErr(code int, msg string) errno.Error {
	return errno.NewError(http.StatusOK, code, msg)
}
