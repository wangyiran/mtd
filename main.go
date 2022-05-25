package main

import (
	"myTodoList/conf"
	"myTodoList/route"
)

func main() {
	//首先加载数据库
	conf.LoadConfig()
	//加载路由
	r := route.NewRouter()
	//启动路由
	r.Run(conf.Port)
}
