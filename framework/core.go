package framework /*框架代码存放于framwork文件夹中，业务代码存放于文件夹之外*/

import "net/http"

//框架核心结构
type Core struct {
}

//初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

//框架核心结构实现handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}
