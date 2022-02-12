package gohttp

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

	DEBUG = "true" // 是否打印日志
)

const (
	DefaultTimeout = 30 // 默认请求超时30秒

	JsonContentType = "application/json; charset=utf-8"   // json 请求类型
	FormContentType = "application/x-www-form-urlencoded" // 表单请求类型

	HeaderVersion   = "X-GOHTTP-VERSION"   // gohttp版本
	HeaderBuildTime = "X-GOHTTP-BUILDTIME" // gohttp编译时间
	HeaderTime      = "X-GOHTTP-TIME"      // 请求时间
	HeaderSign      = "X-GOHTTP-SIGN"      // 签名
	HeaderEncrypt   = "X-GOHTTP-ENCRYPT"   // 参数是否加密 YES表示加密
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
	header[HeaderVersion] = VERSION
	header[HeaderBuildTime] = BUILD_TIME

	cookieJar, _ := cookiejar.New(nil)

	return &GoHttp{
		BaseAddress: "",
		Header:      header,
		Timeout:     DefaultTimeout,
		cookieJar:   cookieJar,
		Client: http.Client{
			Timeout: time.Duration(DefaultTimeout) * time.Second,
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
		timeout = DefaultTimeout
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
