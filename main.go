package main

import (
	"coredemo/framework"
	"net/http"
)

//主函数
func main() {
	server := &http.Server{
		Handler: framework.NewCore(), /*handler字段为空会使用默认DefaultServerMux
		  路由器来填充这个值，一般都使用自定义的请求核心处理函数*/
		Addr: "localhost:8080",
	}
	//通过监听URL地址和控制器函数来创建http服务
	server.ListenAndServe()
}
