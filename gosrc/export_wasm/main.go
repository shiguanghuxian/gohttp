package main

import (
	"encoding/json"
	"gohttp/gosrc/gohttp"
	"log"
	"syscall/js"
)

// 参考 https://withblue.ink/2020/10/03/go-webassembly-http-requests-and-promises.html

var (
	goHttp = gohttp.NewGoHttp()
)

// GetVersion 获取gohttp版本信息
func GetVersion() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return map[string]interface{}{
			"version":    gohttp.VERSION,
			"build_time": gohttp.BUILD_TIME,
			"go_version": gohttp.GO_VERSION,
			"git_hash":   gohttp.GIT_HASH,
		}
	})
}

// SetBaseAddress 设置接口请求根地址
func SetBaseAddress() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			log.Println("未传请求根地址参数")
			return map[string]interface{}{
				"err":      "未传请求根地址参数",
				"err_code": gohttp.CodeFail,
			}
		}
		goHttp.SetBaseAddress(args[0].String())
		return nil
	})
}

// SetTimeout 设置接口请求超时 单位s
func SetTimeout() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			log.Println("未传超时参数")
			return map[string]interface{}{
				"err":      "未传超时参数",
				"err_code": gohttp.CodeFail,
			}
		}
		goHttp.SetTimeout(int64(args[0].Int()))
		return nil
	})
}

// SetHeader 设置公共头
func SetHeader() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			log.Println("请传人两个参数，头下标和值")
			return map[string]interface{}{
				"err":      "请传人两个参数，头下标和值",
				"err_code": gohttp.CodeFail,
			}
		}
		goHttp.SetHeader(args[0].String(), args[1].String())
		return nil
	})
}

// Post 发起post请求
func Post() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return request("POST", args)
	})
}

// Get 发起get请求
func Get() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return request("GET", args)
	})
}

// Put 发起put请求
func Put() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return request("PUT", args)
	})
}

// Delete 发起delete请求
func Delete() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return request("DELETE", args)
	})
}

// Request 发起自定义请求
func Request() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			log.Println("发起自定义请求至少需要两个参数，method和url")
			return map[string]interface{}{
				"err":      "发起自定义请求至少需要两个参数，method和url",
				"err_code": gohttp.CodeFail,
			}
		}
		method := args[0].String()

		return request(method, args[1:])
	})
}

// 具体解析数据调用
func request(method string, args []js.Value) interface{} {
	if len(args) < 1 {
		log.Println("请求至少需要传入url")
		return map[string]interface{}{
			"err":      "请求至少需要传入url",
			"err_code": gohttp.CodeFail,
		}
	}
	url := args[0].String()
	var params map[string]interface{}
	if len(args) > 1 && !args[1].IsNull() && !args[1].IsUndefined() && !args[1].IsNaN() {
		err := json.Unmarshal([]byte(args[1].String()), &params)
		if err != nil {
			log.Println("params参数解析错误", err)
			return map[string]interface{}{
				"err":      "params参数解析错误",
				"err_code": gohttp.CodeFail,
			}
		}
	}
	header := make(map[string]string)
	if len(args) > 2 && !args[2].IsNull() && !args[2].IsUndefined() && !args[2].IsNaN() {
		err := json.Unmarshal([]byte(args[2].String()), &header)
		if err != nil {
			log.Println("header参数解析错误", err)
			return map[string]interface{}{
				"err":      "header参数解析错误",
				"err_code": gohttp.CodeFail,
			}
		}
	}
	var contentType string
	if len(args) > 3 {
		contentType = args[3].String()
	}

	gohttp.Log("请求", url, params, header, contentType)

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0] // 成功返回
		reject := args[1]  // 错误返回
		go func() {
			responseData := goHttp.Request(method, &gohttp.RequestData{
				Url:         url,
				Params:      params,
				Header:      header,
				ContentType: contentType,
			})
			data := map[string]interface{}{
				"err":         responseData.Err,
				"err_code":    responseData.ErrCode,
				"status":      responseData.Status,
				"status_code": responseData.StatusCode,
				"body":        responseData.Body,
				"total_time":  responseData.TotalTime.Milliseconds(),
			}
			// 有错误
			if responseData.ErrCode != 0 {
				reject.Invoke(data)
			} else { // 无错误
				resolve.Invoke(data)
			}
		}()
		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func main() {
	js.Global().Set("GetVersion", GetVersion())
	js.Global().Set("SetBaseAddress", SetBaseAddress())
	js.Global().Set("SetTimeout", SetTimeout())
	js.Global().Set("SetHeader", SetHeader())

	js.Global().Set("Post", Post())
	js.Global().Set("Get", Get())
	js.Global().Set("Put", Put())
	js.Global().Set("Delete", Delete())
	js.Global().Set("Request", Request())

	select {}
}
