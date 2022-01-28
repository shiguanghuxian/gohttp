package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/get", func(rw http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"a": 9090,
		}
		js, _ := json.Marshal(data)
		rw.Write(js)
		rw.WriteHeader(200)
	})
	http.Handle("/", http.FileServer(http.Dir(`./server`)))

	http.ListenAndServe(`:9999`, nil)
}
