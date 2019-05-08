package e

const (
	WARNNING          int    = 10000
	SUCCESS           int    = 20000
	NOCONTENT         int    = 20001
	PARAM_ERROR       int    = 40000
	TOKEN_EMPTY       int    = 50000
	TOKEN_INVAILD     int    = 50001
	TOKEN_EXPRIED     int    = 50002
	ERROR             int    = 30001
	SUCCESS_MSG       string = "成功"
	PARAM_ERROR_MSG   string = "请求参数错误"
	TOKEN_EMPTY_MSG   string = "请求未携带token，无权限访问"
	TOKEN_INVAILD_MSG string = "token非法"
	TOKEN_EXPRIED_MSG string = "token过期"
	ERROR_MSG         string = "未知错误"
)

func GetCodeMessage(code int) string {
	var message string
	switch code {
	case SUCCESS:
		message = SUCCESS_MSG
	case TOKEN_EMPTY:
		message = TOKEN_EMPTY_MSG
	case TOKEN_INVAILD:
		message = TOKEN_INVAILD_MSG
	case TOKEN_EXPRIED:
		message = TOKEN_EXPRIED_MSG
	default:
		message = SUCCESS_MSG
	}
	return message
}
