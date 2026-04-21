package errno

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
