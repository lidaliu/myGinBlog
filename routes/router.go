package routes

import (
	"github.com/gin-gonic/gin"
	"myGinBlog/api/v1"
	"myGinBlog/middleware"
	"myGinBlog/utils"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToke())
	{
		//用户模块路由接口

		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		// 分类模块路由接口
		auth.POST("category/add",v1.AddCategory)
		auth.PUT("category/:id",v1.EditCategory)
		auth.DELETE("category/:id",v1.DeleteCategory)
		// 文章模块路由接口
		auth.POST("article/add",v1.AddArticle)
		auth.PUT("article/:id",v1.EditArticle)
		auth.DELETE("article/:id",v1.DeleteArticle)

	}

	router := r.Group("api/v1")
	{
		//用户模块路由接口
		router.POST("user/add",v1.AddUser)
		router.POST("user/login",v1.Login)
		router.GET("users",v1.GetUsers)
		// 分类模块路由接口
		router.GET("categories",v1.GetCategories)
		// 文章模块路由接口
		router.GET("articles",v1.GetArticles)
		router.GET("article/:id",v1.GetArticleInfo)
		router.GET("article/category/:id",v1.GetArticleByCategory)
	}

	r.Run(utils.HttpPort)
}
