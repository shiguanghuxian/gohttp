package main

import "C"
import (
	"encoding/json"
	"log"
)

/* 导出C函数 */

var (
	goHttp = NewGoHttp()
)

//export GetVersion
func GetVersion() *C.char {
	version := map[string]string{
		"version":    VERSION,
		"build_time": BUILD_TIME,
		"go_version": GO_VERSION,
		"git_hash":   GIT_HASH,
	}
	data, _ := json.Marshal(version)
	return C.CString(string(data))
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
		log.Println("设置请求头，参数解析错误", *in)
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
	log.Println("请求参数", *in)
	// 返回数据
	var responseData *ResponseData
	defer func() {
		data, err := json.Marshal(responseData)
		if err != nil {
			log.Println("响应数据，json编码错误")
		}
		out = C.CString(string(data))
	}()
	// 处理入参
	var requestData *RequestData
	requestData, responseData = strToRequestData(*in)
	if responseData != nil {
		return
	}

	responseData = goHttp.Request("POST", requestData)
	return
}

// 转换入参
func strToRequestData(in string) (*RequestData, *ResponseData) {
	requestData := new(RequestData)
	err := json.Unmarshal([]byte(in), requestData)
	if err != nil {
		log.Println("解析入参错误", in, err)
		return nil, ParameterError
	}
	return requestData, nil
}

func main() {
}
