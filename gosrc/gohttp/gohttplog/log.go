//go:build !wasm
// +build !wasm

package gohttplog

import "log"

func Log(v ...interface{}) {
	if DEBUG {
		return
	}
	// console := js.Global().Get("console")
	// console.Call("log", v...)
	log.Println(v...)
}
