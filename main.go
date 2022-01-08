package main

import (
	"coredemo/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	//通过监听URL地址和控制器函数来创建http服务
	server.ListenAndServe()
}
