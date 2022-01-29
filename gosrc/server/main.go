package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// 模拟数据返回
type Msg struct {
	Code    int32
	Message string
	Data    interface{}
}

func main() {
	http.HandleFunc("/v1/get", get)
	http.HandleFunc("/v1/post", post)
	http.HandleFunc("/v1/put", put)
	http.HandleFunc("/v1/delete", delete)

	http.Handle("/", http.FileServer(http.Dir(`./`)))

	http.ListenAndServe(`:9999`, nil)
}

func get(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		rw.WriteHeader(404)
		return
	}
	time.Sleep(3 * time.Second)
	data := &Msg{
		Code:    1,
		Message: "success",
		Data:    "这是一个get请求",
	}
	js, _ := json.Marshal(data)
	rw.WriteHeader(200)
	rw.Write(js)
}

func post(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.WriteHeader(404)
		return
	}
	data := &Msg{
		Code:    1,
		Message: "success",
		Data:    "这是一个post请求",
	}
	js, _ := json.Marshal(data)
	rw.WriteHeader(200)
	rw.Write(js)
}

func put(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		rw.WriteHeader(404)
		return
	}
	data := &Msg{
		Code:    1,
		Message: "success",
		Data:    "这是一个put请求",
	}
	js, _ := json.Marshal(data)
	rw.WriteHeader(200)
	rw.Write(js)
}

func delete(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		rw.WriteHeader(404)
		return
	}
	data := &Msg{
		Code:    1,
		Message: "success",
		Data:    "这是一个delete请求",
	}
	js, _ := json.Marshal(data)
	rw.WriteHeader(200)
	rw.Write(js)
}
