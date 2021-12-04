package main

import (
	"coredemo/framework"
	"net/http"
)

//主函数
func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
