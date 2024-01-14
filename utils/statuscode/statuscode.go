package statuscode

// nolint
const (
	StatusOK            = 0
	ParamsInvalid       = 40000
	AuthError           = 40001
	PackageNotFound     = 40002
	AppNameRepeat       = 40003
	MetricsNotFound     = 40004
	AppNotFound         = 40005
	AppParamsInvalid    = 40006
	InsecurePassword    = 40007
	ServerInternalError = 50000
	PackageQueryError   = 50001
	MetricsQueryError   = 50002
	UserQueryError      = 50003
)

var codeMsg = map[int]string{
	StatusOK:            "请求处理成功",
	ParamsInvalid:       "参数错误",
	AuthError:           "用户名密码错误",
	PackageNotFound:     "安装包不存在",
	AppNameRepeat:       "应用名称重复",
	MetricsNotFound:     "监控指标不存在",
	AppNotFound:         "应用不存在",
	AppParamsInvalid:    "应用参数校验不合法",
	InsecurePassword:    "不安全密码",
	ServerInternalError: "服务内部错误",
	PackageQueryError:   "查询安装包出错",
	MetricsQueryError:   "指标查询出错",
	UserQueryError:      "查询用户出错",
}

// CodeToMessage covert code to message
func CodeToMessage(code int) string {
	msg, ok := codeMsg[code]
	if !ok {
		return "Unknown"
	}
	return msg
}
