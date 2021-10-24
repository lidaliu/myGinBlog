package routes

import (
	"github.com/gin-gonic/gin"
	"myGinBlog/api/v1"
	"myGinBlog/utils"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.PUT("user/:id",v1.EditUser)
		router.DELETE("user/:id",v1.DeleteUser)
		// 分类模块路由接口
		router.POST("category/add",v1.AddCategory)
		router.GET("categories",v1.GetCategories)
		router.PUT("category/:id",v1.EditCategory)
		router.DELETE("category/:id",v1.DeleteCategory)
		// 文章模块路由接口
		router.POST("article/add",v1.AddArticle)
		router.GET("articles",v1.GetArticles)
		router.GET("article/:id",v1.GetArticleInfo)
		router.GET("article/category/:id",v1.GetArticleByCategory)
		router.PUT("article/:id",v1.EditArticle)
		router.DELETE("article/:id",v1.DeleteArticle)

	}

	r.Run(utils.HttpPort)
}
