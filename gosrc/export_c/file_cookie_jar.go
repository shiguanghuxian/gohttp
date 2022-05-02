package main

import (
	"bytes"
	"encoding/gob"
	"gohttp/gosrc/gohttp/cookiejar"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

/* 是要文件缓存cookie信息 */

type FileCookieJar struct {
	*cookiejar.Jar
	initKeys map[string]bool
	filePath string
}

// 创建一个文件cookie缓存对象
func NewFileCookieJar(filePath string) *FileCookieJar {
	jar, _ := cookiejar.New(nil)
	return &FileCookieJar{
		Jar:      jar,
		initKeys: make(map[string]bool),
		filePath: filePath,
	}
}

func (fcj *FileCookieJar) SaveCookies(u *url.URL) {
	// 文件路径
	cookiePath := fcj.GetCookiePath(u)
	// 获取要存储的原始数据
	cookieData := fcj.GetEntryMap(u)
	if cookieData == nil {
		os.Remove(cookiePath)
		return
	}
	// 序列化
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&cookiejar.CacheCookie{
		URL:      u,
		EntryMap: cookieData,
	})
	if err != nil {
		log.Println("序列号cookie map错误", err, cookieData)
		return
	}
	// 保存
	err = ioutil.WriteFile(cookiePath, buffer.Bytes(), fs.ModePerm)
	if err != nil {
		log.Println("cookie写文件错误", err, cookieData)
		return
	}
}

func (fcj *FileCookieJar) InitCookies(u *url.URL) {
	if fcj.initKeys[fcj.GetJarKey(u)] {
		return
	}
	fcj.initKeys[fcj.GetJarKey(u)] = true
	// 文件路径
	cookiePath := fcj.GetCookiePath(u)
	// 从文件读取
	cookieJs, err := ioutil.ReadFile(cookiePath)
	if err != nil {
		log.Println("cookie读文件错误", err, cookiePath)
		return
	}
	if len(cookieJs) == 0 {
		return
	}
	cookieData := new(cookiejar.CacheCookie)
	decoder := gob.NewDecoder(bytes.NewReader(cookieJs))
	err = decoder.Decode(&cookieData)
	if err != nil {
		log.Println("cookie缓存gob解析错误", err, cookiePath)
		return
	}
	if cookieData.URL == nil || cookieData.EntryMap == nil {
		log.Println("缓存的cookie为空")
		return
	}
	// 设置cookie
	cookies := make([]*http.Cookie, 0)
	for _, v := range cookieData.EntryMap {
		var sameSite http.SameSite
		switch v.SameSite {
		case "SameSite":
			sameSite = http.SameSiteDefaultMode
		case "SameSite=Strict":
			sameSite = http.SameSiteStrictMode
		case "SameSite=Lax":
			sameSite = http.SameSiteLaxMode
		}
		cookies = append(cookies, &http.Cookie{
			Name:     v.Name,
			Value:    v.Value,
			Path:     v.Path,
			Domain:   v.Domain,
			Expires:  v.Expires,
			Secure:   v.Secure,
			HttpOnly: v.HttpOnly,
			SameSite: sameSite,
		})
	}
	fcj.SetCookies(cookieData.URL, cookies)
}

func (fcj *FileCookieJar) GetCookiePath(u *url.URL) string {
	return filepath.Join(fcj.filePath, fcj.GetJarKey(u))
}
