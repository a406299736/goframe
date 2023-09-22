package code

import "github.com/a406299736/goframe/pkg/errors"

var zhCNText = map[int]string{
	Success:            "SUCCESS",
	ServerError:        "内部服务错误",
	TooManyRequests:    "请求频繁",
	ParamError:         "参数错误",
	AuthorizationError: "授权错误",
	SignError:          "签名错误",
	CacheSetError:      "设置缓存失败",
	CacheGetError:      "获取缓存失败",
	CacheDelError:      "删除缓存失败",
	CacheNotExist:      "缓存不存在",
	ResubmitError:      "请勿重复提交",
	EncodeError:        "加密失败",
	DecodeError:        "解密失败",
	RedisConnectError:  "Redis 连接失败",
	MySQLConnectError:  "MySQL 连接失败",
	MySQLExecError:     "SQL 执行失败",
	ValueIsNil:         "nil pointer",

	NotAllowed:    "无权限",
	UserNotExists: "用户不存在",
	NotExists:     "不存在",

	ThirdRespError: "三方接口返回错误",

	ParseError:   "解析失败",
	KeyNotExists: "key不存在",
}

var NilEr = errors.NewErr(ValueIsNil, zhCNText[ValueIsNil])                   // nil错误
var ParseEr = errors.NewErr(ParseError, zhCNText[ParamError])                 // 解析错误
var ParamsEr = errors.NewErr(ParamError, zhCNText[ParamError])                // 参数错误
var NotAllowedEr = errors.NewErr(NotAllowed, zhCNText[NotAllowed])            // 无权限
var NotExistEr = errors.NewErr(NotExists, zhCNText[NotExists])                // 不存在
var ThirdEr = errors.NewErr(ThirdRespError, zhCNText[ThirdRespError])         // 三方错误
var AuthEr = errors.NewErr(AuthorizationError, zhCNText[AuthorizationError])  // 授权错误
var ManyRequestEr = errors.NewErr(TooManyRequests, zhCNText[TooManyRequests]) // 请求频繁
var SignEr = errors.NewErr(SignError, zhCNText[SignError])                    // 签名错误
