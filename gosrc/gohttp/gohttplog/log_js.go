//go:build wasm
// +build wasm

package gohttplog

import (
	"fmt"
	"syscall/js"
)

func Log(v ...interface{}) {
	if DEBUG {
		return
	}
	if len(v) == 0 {
		return
	}
	console := js.Global().Get("console")
	// console对象是否获取成功
	if console.IsUndefined() {
		return
	}
	// js防止类型不支持，转为字符串
	vList := make([]interface{}, 0)
	for _, vl := range v {
		vList = append(vList, fmt.Sprint(vl))
	}
	console.Call("log", vList...)
}
