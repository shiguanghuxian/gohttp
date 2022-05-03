package cookiejar

import (
	"net/url"
)

/* 默认cookie管理 */

// 用于缓存数据使用
type CacheCookie struct {
	URL      *url.URL
	EntryMap map[string]Entry
}

func NewDefaultCookieJar() *DefaultCookieJar {
	jar, _ := New(nil)
	return &DefaultCookieJar{
		Jar: jar,
	}
}

type DefaultCookieJar struct {
	*Jar
}

func (dcj *DefaultCookieJar) SaveCookies(u *url.URL) {
	// log.Println("存储cookie", u.String())
}

func (dcj *DefaultCookieJar) InitCookies(u *url.URL) {
	// log.Println("初始化cookie", u.String())
}
