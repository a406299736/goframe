package configs

const (
	// ProjectVersion 项目版本
	ProjectVersion = "v1.2.7"

	// ProjectName 项目名称
	ProjectName = "go-basic"

	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"

	// ProjectPort 项目端口
	ProjectPort = ":8006"

	// HeaderLoginToken 登录验证 Token，Header 中传递的参数
	HeaderLoginToken = "Token"

	// HeaderSignToken 签名验证 Token，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenDate 签名验证 Date，Header 中传递的参数
	HeaderSignTokenDate = "Authorization-Date"

	// RedisKeyPrefixRequestID Redis Key 前缀 - 防止重复提交
	RedisKeyPrefixRequestID = ProjectName + ":request-id:"

	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"

	// RedisKeyPrefixSignature Redis Key 前缀 - 签名验证信息
	RedisKeyPrefixSignature = ProjectName + ":signature:"

	// 简体中文
	ZhCN = "zh-cn"

	// 英文
	EnUS = "en-us"
)
