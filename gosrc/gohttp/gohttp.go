package gohttp

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

/* go实现的http请求库 */

var (
	VERSION    = "0.0.0"
	BUILD_TIME = ""
	GO_VERSION = ""
	GIT_HASH   = ""

	DEBUG = "false" // 是否打印日志

	RATE_LIMIT       = "true" // 是否开启限流，默认开启
	RATE_LIMIT_LIMIT = "15"   // 限流每秒产生token数
	RATE_LIMIT_BURST = "15"   // 限流桶大小
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

	MaxRateLimitLimit  = 20                 // 限流每秒token数最大值，设置值大于此值或小于1时，设置为此值
	MaxRateLimitBurst  = 10                 // 限流桶最大值，设置值大于此值或小于1时，设置为此值
	RateLimitTypeWait  = "wait"             // 当达到上限时等待
	RateLimitTypeAllow = "allow"            // 当达到上限时直接返回
	RateLimitType      = RateLimitTypeAllow // 限流处理类型，根据使用，修改此值
)

// TODO cookie的持久化

type GoHttp struct {
	BaseAddress string            // 请求接口根地址，当接口请求不以http或https开头时使用
	Header      map[string]string // 公共头
	Timeout     int64             // 请求超时 单位秒
	Client      http.Client       // http 客户端
	cookieJar   *cookiejar.Jar    // cookie管理
	limiter     *rate.Limiter     // 限流器
}

func NewGoHttp() *GoHttp {
	header := make(map[string]string)
	header[HeaderVersion] = VERSION
	header[HeaderBuildTime] = BUILD_TIME

	cookieJar, _ := cookiejar.New(nil)

	// 初始化限流器
	var limiter *rate.Limiter
	if isRateLimit() {
		limiter = rate.NewLimiter(getRateLimitLimit(), getRateLimitBurst())
	}

	return &GoHttp{
		BaseAddress: "",
		Header:      header,
		Timeout:     DefaultTimeout,
		cookieJar:   cookieJar,
		Client: http.Client{
			Timeout: time.Duration(DefaultTimeout) * time.Second,
		},
		limiter: limiter,
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

// 请求限流
func (gh *GoHttp) requestLimiter() *ResponseData {
	if gh.limiter == nil {
		return nil
	}
	switch RateLimitType {
	case RateLimitTypeWait:
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(gh.Timeout)*time.Second)
		defer cancel()
		err := gh.limiter.Wait(ctx)
		if err != nil {
			Log("限流等待错误", err)
			return LimiterError
		}
	case RateLimitTypeAllow:
		ok := gh.limiter.Allow()
		if !ok {
			Log("限流Allow")
			return LimiterError
		}
	}
	return nil
}

// 是否开启限流
func isRateLimit() bool {
	return RATE_LIMIT == "true"
}

// 获取限流桶大小
func getRateLimitBurst() int {
	num, err := strconv.Atoi(RATE_LIMIT_BURST)
	if err != nil {
		Log("限流设置错误", err)
	}
	if num < 1 || num > MaxRateLimitBurst {
		num = MaxRateLimitBurst
	}
	return num
}

// 获取限流每秒产生token数
func getRateLimitLimit() rate.Limit {
	num, err := strconv.Atoi(RATE_LIMIT_LIMIT)
	if err != nil {
		Log("限流设置错误", err)
	}
	var r rate.Limit
	if num < 1 || num > MaxRateLimitLimit {
		r = MaxRateLimitLimit
	} else {
		r = rate.Limit(num)
	}
	return r
}
