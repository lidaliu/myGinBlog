package main

import (
	"myGinBlog/model"
	"myGinBlog/routes"
)

func main()  {
	//引用数据库
	model.InitDb()
	//路由
	routes.InitRouter()
}