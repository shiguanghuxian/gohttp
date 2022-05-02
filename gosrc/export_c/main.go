package main

import "C"
import (
	"encoding/json"
	"fmt"
	"gohttp/gosrc/gohttp"
)

/* 导出C函数 */

var (
	goHttp = gohttp.NewGoHttp()
)

//export GetVersion
func GetVersion() *C.char {
	version := map[string]string{
		"version":    gohttp.VERSION,
		"build_time": gohttp.BUILD_TIME,
		"go_version": gohttp.GO_VERSION,
		"git_hash":   gohttp.GIT_HASH,
	}
	data, _ := json.Marshal(version)
	return C.CString(string(data))
}

//export SetCookiePath
func SetCookiePath(in *string) {
	if *in == "" {
		gohttp.Log("cookie存储路径设置为空", *in)
		return
	}
	goHttp.SetCookieJar(NewFileCookieJar(*in))
}

//export SetBaseAddress
func SetBaseAddress(in *string) {
	goHttp.SetBaseAddress(*in)
}

//export SetTimeout
func SetTimeout(in int64) {
	goHttp.SetTimeout(in)
}

//export SetHeader
func SetHeader(in *string) {
	header := make(map[string]string)
	err := json.Unmarshal([]byte(*in), &header)
	if err != nil {
		gohttp.Log("设置请求头，参数解析错误", *in)
		return
	}
	for k, v := range header {
		goHttp.SetHeader(k, v)
	}
}

//export Post
func Post(in *string) *C.char {
	return request("POST", in)
}

//export Get
func Get(in *string) *C.char {
	return request("GET", in)
}

//export Put
func Put(in *string) *C.char {
	return request("PUT", in)
}

//export Delete
func Delete(in *string) *C.char {
	return request("DELETE", in)
}

//export Request
func Request(method *string, in *string) *C.char {
	return request(*method, in)
}

// 具体解析数据调用
func request(method string, in *string) (out *C.char) {
	gohttp.Log("请求参数", *in)
	// 返回数据
	var responseData *gohttp.ResponseData
	defer func() {
		data, err := json.Marshal(responseData)
		if err != nil {
			gohttp.Log("响应数据，json编码错误")
		}
		out = C.CString(string(data))
	}()
	// 处理入参
	var requestData *gohttp.RequestData
	requestData, responseData = strToRequestData(*in)
	if responseData != nil {
		return
	}

	responseData = goHttp.Request(method, requestData)
	return
}

// 转换入参
func strToRequestData(in string) (*gohttp.RequestData, *gohttp.ResponseData) {
	requestData := new(gohttp.RequestData)
	err := json.Unmarshal([]byte(in), requestData)
	if err != nil {
		gohttp.Log("解析入参错误", in, err)
		return nil, gohttp.ParameterError
	}
	return requestData, nil
}

func main() {
	if gohttp.DEBUG == "true" {
		version := map[string]string{
			"version":    gohttp.VERSION,
			"build_time": gohttp.BUILD_TIME,
			"go_version": gohttp.GO_VERSION,
			"git_hash":   gohttp.GIT_HASH,
		}
		data, _ := json.Marshal(version)
		fmt.Println(string(data))
	} else {
		fmt.Println("release")
	}
}
