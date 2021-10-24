package v1

import (
	"github.com/gin-gonic/gin"
	"myGinBlog/model"
	"myGinBlog/utils/errmsg"
	"net/http"
	"strconv"
)
var code int
//查询用户是否存在
func UserExists(c *gin.Context)  {
	
}
// 添加用户
func AddUser(c *gin.Context)  {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum,_ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0{
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})

}
//编辑用户
func EditUser(c *gin.Context)  {
    id,_ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUserExceptSelf(uint(id), data.Username)
	if code == errmsg.SUCCESS{
		code = model.CheckUserById(id)
		if code == errmsg.SUCCESS{
			code = model.EditUser(id,&data)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message": errmsg.GetErrMsg(code),
	})
}
//删除用户
func DeleteUser(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"message": errmsg.GetErrMsg(code),
	})
}


