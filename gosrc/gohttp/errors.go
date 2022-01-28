package gohttp

/* 错误 */

const (
	CodeFail              = -1
	CodeMethodError       = 80000
	CodeParameterError    = 80001
	CodeResponseReadError = 80002
)

var (
	ParameterError = &ResponseData{
		Err:     "参数错误",
		ErrCode: CodeParameterError,
	}
	ResponseReadError = &ResponseData{
		Err:     "读取响应body错误",
		ErrCode: CodeResponseReadError,
	}
	MethodError = &ResponseData{
		Err:     "不支持的请求类型",
		ErrCode: CodeMethodError,
	}
)
