package main

import (
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

/* go实现的http请求库 */

var (
	VERSION    = "0.0.0"
	BUILD_TIME = ""
	GO_VERSION = ""
	GIT_HASH   = ""
)

const (
	DefaultSendTimeout = 30

	JsonContentType = "application/json; charset=utf-8"
)

// TODO cookie的持久化

type GoHttp struct {
	BaseAddress string            // 请求接口根地址，当接口请求不以http或https开头时使用
	Header      map[string]string // 公共头
	Timeout     int64             // 请求超时 单位秒
	Client      http.Client       // http 客户端
	cookieJar   *cookiejar.Jar    // cookie管理
}

func NewGoHttp() *GoHttp {
	header := make(map[string]string)
	header["GOHTTP-Version"] = VERSION
	header["GOHTTP-BuildTime"] = BUILD_TIME

	cookieJar, _ := cookiejar.New(nil)

	return &GoHttp{
		BaseAddress: "",
		Header:      header,
		Timeout:     DefaultSendTimeout,
		cookieJar:   cookieJar,
		Client: http.Client{
			Timeout: time.Duration(DefaultSendTimeout) * time.Second,
		},
	}
}

// 设置头信息
func (gh *GoHttp) SetHeader(name, value string) {
	gh.Header[name] = value
}

// 设置根地址
func (gh *GoHttp) SetBaseAddress(baseAddress string) {
	gh.BaseAddress = baseAddress
}

// 设置请求超时
func (gh *GoHttp) SetTimeout(timeout int64) {
	if timeout <= 0 {
		timeout = DefaultSendTimeout
	}
	gh.Timeout = timeout
	gh.Client.Timeout = time.Duration(timeout) * time.Second
}

// 获取请求完整地址
func (gh *GoHttp) getFullUrl(inUrl string) string {
	if strings.HasPrefix(inUrl, "http://") || strings.HasPrefix(inUrl, "https://") {
		return inUrl
	}
	return gh.BaseAddress + inUrl
}
