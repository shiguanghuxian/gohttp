package gohttp

import "log"

func Log(v ...interface{}) {
	if DEBUG != "true" {
		return
	}
	log.Println(v...)
}
