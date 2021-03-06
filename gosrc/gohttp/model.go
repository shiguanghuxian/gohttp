package gohttp

import "time"

/* 数据模型 */

// 导出c函数请求参数
type RequestData struct {
	Url         string                 `json:"url"`          // 请求地址
	Params      map[string]interface{} `json:"params"`       // get 或 post 参数
	Header      map[string]string      `json:"header"`       // 请求头
	ContentType string                 `json:"content_type"` // post或put请求body类型
	Encrypt     bool                   `json:"encrypt"`      // 是否加密请求 服务端限制某些接口必须加密请求
}

// 导出c函数响应json字符串
type ResponseData struct {
	Err        string        `json:"err"`         // 是否遇到错误
	ErrCode    int32         `json:"err_code"`    // 错误代码
	Status     string        `json:"status"`      // e.g. "200 OK"
	StatusCode int           `json:"status_code"` // e.g. 200
	Body       string        `json:"body"`        // 响应内容
	TotalTime  time.Duration `json:"total_time"`  // 总耗时 包括go数据处理和http请求耗时
}
