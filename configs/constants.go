package configs

const (
	ProjectName = "go-basic"
	ProjectPort = ":8006"

	HeaderLoginToken = "Token"
	HeaderSignToken = "Authorization"
	HeaderSignTokenDate = "Authorization-Date"
	RedisKeyPrefixRequestID = ProjectName + ":request-id:"
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"

	ZhCN = "zh-cn"
)
