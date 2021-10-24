package v1

import (
	"github.com/gin-gonic/gin"
	"myGinBlog/model"
	"myGinBlog/utils/errmsg"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context)  {
	var data model.Article
	_ =  c.ShouldBindJSON(&data)
	model.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data": &data,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询单个文章信息
func GetArticleInfo(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message": errmsg.GetErrMsg(code),
	})
}
//根据分类查询分类下的文章
func GetArticleByCategory(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	pageSize,_:= strconv.Atoi(c.Param("pageSize"))
	pageNum,_:= strconv.Atoi(c.Param("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data,code := model.GetArticleByCategory(pageSize,pageNum,id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticles(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum ,_ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data,code := model.GetArticles(pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
//编辑文章
func EditArticle(c *gin.Context)  {
	var data  model.Article
	id,_ := strconv.Atoi(c.Param("id"))

	_=c.ShouldBindJSON(&data)
	if code == errmsg.SUCCESS{
		code = model.CheckArticleById(id)
		if code == errmsg.SUCCESS{
			code = model.EditArticle(id,&data)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})


}
//删除分类

func DeleteArticle(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message":errmsg.GetErrMsg(code),
	})


}
