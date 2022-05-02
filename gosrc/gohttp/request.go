package gohttp

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/wumansgy/goEncrypt"
)

/* 几种http请求 */

func (gh *GoHttp) Request(method string, requestData *RequestData) (responseData *ResponseData) {
	Log("开始处理请求", requestData.Url)
	defer Log("结束处理请求", requestData.Url)

	startTime := time.Now()
	defer func() {
		if responseData != nil {
			responseData.TotalTime = time.Since(startTime)
		}
	}()

	// 限流
	limitErr := gh.requestLimiter()
	if limitErr != nil {
		responseData = limitErr
		return
	}

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
		Log("不支持的请求类型")
		return MethodError
	}
	if responseData != nil {
		return
	}

	// 签名
	signT, signStr := gh.signature(reqUrl, method, requestData)

	// 参数加密
	if requestData.Encrypt {
		switch method {
		case "POST", "PUT":
			body = gh.encryptBodyParams(body, signStr)
		case "GET", "DELETE":
			reqUrl = gh.encryptUrlParams(reqUrl, signStr)
		}
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

	// 签名头
	req.Header.Set(HeaderTime, signT)
	req.Header.Set(HeaderSign, signStr)

	// 标记加密
	if requestData.Encrypt {
		req.Header.Set(HeaderEncrypt, "YES")
	}

	js, _ := json.Marshal(req.Header)
	Log(string(js))

	// 尝试初始化cookie
	gh.cookieJar.InitCookies(req.URL)
	// 记录请求前后cookie变化
	cookieOldHash := gh.cookieJar.GetHash(req.URL)
	defer func() {
		// 发生变化重新保存cookie到文件
		if cookieOldHash != gh.cookieJar.GetHash(req.URL) {
			gh.cookieJar.SaveCookies(req.URL)
			Log("发生cookie变化", gh.cookieJar.Cookies(req.URL))
		}
	}()

	Log("请求cookie", gh.cookieJar.Cookies(req.URL))

	resp, err := gh.Client.Do(req)
	if err != nil {
		Log(err)
		return &ResponseData{
			Err:     err.Error(),
			ErrCode: CodeFail,
		}
	}
	Log("响应cookie", resp.Cookies())
	return gh.ResponseToResponseData(resp)
}

// post请求
func (gh *GoHttp) POST(requestData *RequestData) (contentType string, body io.Reader, responseData *ResponseData) {
	if gh.isJsonRequest(requestData) {
		contentType = JsonContentType
		body = gh.requestParamsToJsonStr(requestData.Params)
	} else {
		contentType = FormContentType
		body = gh.requestParamsToFormStr(requestData.Params)
	}
	return
}

// get请求
func (gh *GoHttp) GET(requestData *RequestData) (string, *ResponseData) {
	urlInfo, err := url.Parse(gh.getFullUrl(requestData.Url))
	if err != nil {
		Log("解析get请求url格式错误", requestData.Url, err)
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
		contentType = FormContentType
		body = gh.requestParamsToFormStr(requestData.Params)
	}
	return
}

// delete请求
func (gh *GoHttp) DELETE(requestData *RequestData) (string, *ResponseData) {
	urlInfo, err := url.Parse(gh.getFullUrl(requestData.Url))
	if err != nil {
		Log("解析get请求url格式错误", requestData.Url, err)
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
		Log(err)
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
		Log("请求参数转json错误", err)
	}
	return bytes.NewReader(body)
}

// post或put form请求body
func (gh *GoHttp) requestParamsToFormStr(params map[string]interface{}) io.Reader {
	if len(params) == 0 {
		return bytes.NewReader(nil)
	}
	data := url.Values{}
	for k, v := range params {
		data.Set(k, fmt.Sprint(v))
	}
	return strings.NewReader(data.Encode())
}

// 计算签名，返回计算后的签名和头信息
func (gh *GoHttp) signature(reqUrl string, method string, requestData *RequestData) (t string, sign string) {
	params := make(map[string]string)
	if method == "POST" || method == "PUT" {
		for k, v := range requestData.Params {
			params[k] = fmt.Sprint(v)
		}
	} else {
		urlInfo, err := url.Parse(gh.getFullUrl(reqUrl))
		if err != nil {
			Log("签名计算解析get请求url格式错误", reqUrl, err)
		}
		queryParams := urlInfo.Query()
		for k, v := range queryParams {
			if len(v) == 0 {
				continue
			}
			params[k] = v[0]
		}
	}
	// 16进制时间戳
	t = strings.ToUpper(fmt.Sprintf("%x", time.Now().Unix()))
	sign = Signature(params, t)
	return
}

// 加密get或delete的url参数
func (gh *GoHttp) encryptUrlParams(reqUrl string, sign string) string {
	urlInfo, err := url.Parse(gh.getFullUrl(reqUrl))
	if err != nil {
		Log("加密，解析get请求url格式错误", reqUrl, err)
		return reqUrl
	}
	q, err := goEncrypt.AesCbcEncrypt([]byte(urlInfo.RawQuery), GetEncryptKey(sign), GetEncryptIv(sign))
	if err != nil {
		Log("aes加密错误url", reqUrl, sign, err)
		return reqUrl
	}
	// 只保留q参数
	data := url.Values{}
	data.Set("q", hex.EncodeToString(q))
	urlInfo.RawQuery = data.Encode()
	return urlInfo.String()
}

// 加密请求体
func (gh *GoHttp) encryptBodyParams(body io.Reader, sign string) io.Reader {
	if body == nil {
		return nil
	}
	bodyVal, err := ioutil.ReadAll(body)
	if err != nil {
		Log("加密读取body错误", err)
		return nil
	}
	q, err := goEncrypt.AesCbcEncrypt(bodyVal, GetEncryptKey(sign), GetEncryptIv(sign))
	if err != nil {
		Log("aes加密错误body", sign, err)
		return bytes.NewBuffer(bodyVal)
	}
	return bytes.NewBuffer(q)
}
