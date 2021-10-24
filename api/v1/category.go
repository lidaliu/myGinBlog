package v1

import (
	"github.com/gin-gonic/gin"
	"myGinBlog/model"
	"myGinBlog/utils/errmsg"
	"net/http"
	"strconv"
)
//查询分类名是否存在

//添加分类
func AddCategory(c *gin.Context)  {
	var data model.Category
	_ =  c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS{
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data": &data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询某个分类下的文章

//查询分类列表
func GetCategories(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum ,_ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategories(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
//编辑分类
func EditCategory(c *gin.Context)  {
	var data  model.Category
	id,_ := strconv.Atoi(c.Param("id"))

	_=c.ShouldBindJSON(&data)
	code = model.CheckCategoryExceptSelf(uint(id), data.Name)
	if code == errmsg.SUCCESS{
		code = model.CheckCategoryById(id)
		if code == errmsg.SUCCESS{
			code = model.EditCategory(id,&data)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})


}
//删除分类

func DeleteCategory(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCategory(id)
    c.JSON(http.StatusOK,gin.H{
    	"status": code,
    	"message":errmsg.GetErrMsg(code),
	})


}