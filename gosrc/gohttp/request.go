package gohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/* 几种http请求 */

func (gh *GoHttp) Request(method string, requestData *RequestData) (responseData *ResponseData) {
	log.Println("开始处理请求", requestData.Url)
	defer log.Println("结束处理请求", requestData.Url)

	method = strings.ToUpper(method)
	reqUrl := gh.getFullUrl(requestData.Url)
	var body io.Reader
	var contentType string
	switch method {
	case "POST":
		contentType, body, responseData = gh.POST(requestData)
	case "GET":
		reqUrl, responseData = gh.GET(requestData)
	case "PUT":
		contentType, body, responseData = gh.PUT(requestData)
	case "DELETE":
		reqUrl, responseData = gh.DELETE(requestData)
	default:
		log.Println("不支持的请求类型")
		return MethodError
	}
	if responseData != nil {
		return
	}
	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return &ResponseData{
			Err:     err.Error(),
			ErrCode: CodeFail,
		}
	}
	// 公共头
	for k, v := range gh.Header {
		req.Header.Set(k, v)
	}

	// 本次请求头
	for k, v := range requestData.Header {
		req.Header.Set(k, v)
	}

	// 请求内容类型
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	js, _ := json.Marshal(req.Header)
	log.Println(string(js))

	resp, err := gh.Client.Do(req)
	if err != nil {
		log.Println(err)
		return &ResponseData{
			Err:     err.Error(),
			ErrCode: CodeFail,
		}
	}
	return gh.ResponseToResponseData(resp)
}

// post请求
func (gh *GoHttp) POST(requestData *RequestData) (contentType string, body io.Reader, responseData *ResponseData) {
	if gh.isJsonRequest(requestData) {
		contentType = JsonContentType
		body = gh.requestParamsToJsonStr(requestData.Params)
	} else {
		data := url.Values{}
		for k, v := range requestData.Params {
			data.Set(k, fmt.Sprint(v))
		}
		body = strings.NewReader(data.Encode())
		contentType = FormContentType
	}
	return
}

// get请求
func (gh *GoHttp) GET(requestData *RequestData) (string, *ResponseData) {
	urlInfo, err := url.Parse(gh.getFullUrl(requestData.Url))
	if err != nil {
		log.Println("解析get请求url格式错误", requestData.Url, err)
		return "", &ResponseData{
			Err:     err.Error(),
			ErrCode: CodeFail,
		}
	}
	// 添加请求参数
	query := urlInfo.Query()
	for k, v := range requestData.Params {
		query.Set(k, fmt.Sprint(v))
	}
	urlInfo.RawQuery = query.Encode()

	return urlInfo.String(), nil
}

// put请求
func (gh *GoHttp) PUT(requestData *RequestData) (contentType string, body io.Reader, responseData *ResponseData) {
	if gh.isJsonRequest(requestData) {
		contentType = JsonContentType
		body = gh.requestParamsToJsonStr(requestData.Params)
	} else {
		data := url.Values{}
		for k, v := range requestData.Params {
			data.Set(k, fmt.Sprint(v))
		}
		body = strings.NewReader(data.Encode())
		contentType = FormContentType
	}
	return
}

// delete请求
func (gh *GoHttp) DELETE(requestData *RequestData) (string, *ResponseData) {
	urlInfo, err := url.Parse(gh.getFullUrl(requestData.Url))
	if err != nil {
		log.Println("解析get请求url格式错误", requestData.Url, err)
		return "", &ResponseData{
			Err:     err.Error(),
			ErrCode: CodeFail,
		}
	}
	// 添加请求参数
	query := urlInfo.Query()
	for k, v := range requestData.Params {
		query.Set(k, fmt.Sprint(v))
	}
	urlInfo.RawQuery = query.Encode()

	return urlInfo.String(), nil
}

// 转换数据类型为c函数导出结构
func (gh *GoHttp) ResponseToResponseData(resp *http.Response) *ResponseData {
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ResponseReadError
	}
	return &ResponseData{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}
}

// 判断请求是否以json形式
func (gh *GoHttp) isJsonRequest(requestData *RequestData) bool {
	if requestData != nil && strings.Contains(strings.ToLower(requestData.ContentType), "application/json") {
		return true
	}
	for k, v := range gh.Header {
		if strings.ToLower(k) == "content-type" {
			return strings.Contains(strings.ToLower(v), "application/json")
		}
	}
	return false
}

// post或put json请求body
func (gh *GoHttp) requestParamsToJsonStr(params map[string]interface{}) io.Reader {
	if len(params) == 0 {
		return bytes.NewReader(nil)
	}
	body, err := json.Marshal(params)
	if err != nil {
		log.Println("请求参数转json错误", err)
	}
	return bytes.NewReader(body)
}
