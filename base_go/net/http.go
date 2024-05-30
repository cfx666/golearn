package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/test", func(rsp http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			fmt.Fprintln(rsp, "Hello World")
		default:
			rsp.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(rsp, "Method Not Allowed")
		}

	})

	http.ListenAndServe(":8080", nil)
}
