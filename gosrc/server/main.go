package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gohttp/gosrc/gohttp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt"
)

// 模拟数据返回
type Msg struct {
	Code    int32
	Message string
	Data    interface{}
}

func main() {
	ginEngine := gin.Default()
	ginEngine.Static("/demo", "./")

	ginEngine.Use(decryptParams())  // 参数加密
	ginEngine.Use(checkSignature()) // 签名
	ginEngine.GET("/v1/get", get)
	ginEngine.POST("/v1/post", post)
	ginEngine.PUT("/v1/put", put)
	ginEngine.DELETE("/v1/delete", delete)

	ginEngine.Run(":9999")
}

// 验证接口签名
func checkSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		signT := c.GetHeader(gohttp.HeaderTime)
		signStr := c.GetHeader(gohttp.HeaderSign)

		params := make(map[string]string)

		switch c.Request.Method {
		case "POST", "PUT":
			var bodyBytes []byte
			bodyParams := make(map[string]interface{}, 0)
			if c.Request.Body != nil {
				var err error
				bodyBytes, err = ioutil.ReadAll(c.Request.Body)
				if err != nil {
					log.Println("读取body错误", string(bodyBytes), err)
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}

				if strings.Contains(strings.ToLower(c.GetHeader("Content-Type")), "application/json") {
					err = json.Unmarshal(bodyBytes, &bodyParams)
				} else {
					var urlValues url.Values
					urlValues, err = url.ParseQuery(string(bodyBytes))
					for k, v := range urlValues {
						if len(v) < 1 {
							continue
						}
						bodyParams[k] = v[0]
					}
				}
				if err != nil {
					log.Println("解析请求体错误", string(bodyBytes), err)
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			// 转换为string map
			for k, v := range bodyParams {
				params[k] = fmt.Sprint(v)
			}
		case "GET", "DELETE":
			urlQuery := c.Request.URL.Query()
			for k, v := range urlQuery {
				if len(v) == 0 {
					continue
				}
				params[k] = v[0]
			}
		}
		// 验证签名
		sign := gohttp.Signature(params, signT)
		if sign != signStr {
			log.Println("签名校验错误", signT, sign, signStr)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// 解密加密请求
func decryptParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		isEncrypt := c.GetHeader(gohttp.HeaderEncrypt)
		if strings.ToUpper(isEncrypt) != "YES" {
			return
		}
		signStr := c.GetHeader(gohttp.HeaderSign)
		if len(signStr) != 32 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var cipherText []byte // 加密串
		var err error
		switch c.Request.Method {
		case "POST", "PUT":
			if c.Request.Body != nil {
				cipherText, err = ioutil.ReadAll(c.Request.Body)
				if err != nil {
					log.Println("读取请求体错误", err)
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
		case "GET", "DELETE":
			qStr := c.Query("q")
			if qStr != "" {
				cipherText, err = hex.DecodeString(qStr)
				if err != nil {
					log.Println("读取url参数错误", err)
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
		}
		if len(cipherText) > 0 {
			plaintext, err := goEncrypt.AesCbcDecrypt(cipherText, gohttp.GetEncryptKey(signStr), gohttp.GetEncryptIv(signStr))
			if err != nil {
				log.Println("aes解密参数错误", cipherText, err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			switch c.Request.Method {
			case "POST", "PUT":
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(plaintext))
			case "GET", "DELETE":
				c.Request.URL.RawQuery = string(plaintext)
			}
		}
	}
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, &Msg{
		Code:    1,
		Message: "success",
		Data:    c.Request.URL.RawQuery,
	})
}

func post(c *gin.Context) {
	var data interface{}
	err := c.Bind(&data)
	if err != nil {
		log.Println("post解析参数错误", err)
	}
	c.JSON(http.StatusOK, &Msg{
		Code:    1,
		Message: "success",
		Data:    data,
	})
}

func put(c *gin.Context) {
	var data interface{}
	err := c.ShouldBindQuery(&data)
	if err != nil {
		log.Println("post解析参数错误", err)
	}
	c.JSON(http.StatusOK, &Msg{
		Code:    1,
		Message: "success",
		Data:    data,
	})
}

func delete(c *gin.Context) {
	c.JSON(http.StatusOK, &Msg{
		Code:    1,
		Message: "success",
		Data:    c.Request.URL.RawQuery,
	})
}
