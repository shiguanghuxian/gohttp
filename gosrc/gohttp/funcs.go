package gohttp

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// sha256 但只返回32位
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	str := hex.EncodeToString(h.Sum(nil))
	start, _ := strconv.Atoi(str[:1])
	Log("hash剪切起始位置", start)
	return str[start : 32+start]
}

// md5
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 计算签名
func Signature(params map[string]string, t string) (sign string) {
	keyList := make([]string, 0)
	for k := range params {
		keyList = append(keyList, k)
	}

	// 排序一下
	sort.Strings(keyList)

	str := ""
	for _, v := range keyList {
		if str != "" {
			str += "&"
		}
		str += fmt.Sprintf("%s=%s", v, params[v])
	}

	Log("签名", str, VERSION, t)
	// 计算签名
	sign = genSignature(t, str)
	return
}

// 具体计算签名 可修改次函数实现自定义签名算法
func genSignature(t string, in string) string {
	t16 := strings.ToLower(t)
	return HashString(HashString(in) + HashString(t+VERSION) + t16)
}

// 参数加密key
func GetEncryptKey(sign string) []byte {
	return []byte(GetMd5String(sign))
}

// 参数加密向量
func GetEncryptIv(sign string) []byte {
	return []byte(sign[:16])
}
