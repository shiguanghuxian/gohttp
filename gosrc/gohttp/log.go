package gohttp

import "log"

func Log(v ...interface{}) {
	if DEBUG != "true" {
		return
	}
	// console := js.Global().Get("console")
	// console.Call("log", v...)
	log.Println(v...)
}
