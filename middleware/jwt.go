package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myGinBlog/utils"
	"myGinBlog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string)(string,int)  {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
			Issuer: "myGinBlog",
		},
	}

	reqClaim :=jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token,err:=reqClaim.SignedString(JwtKey)
	if err != nil {
		return "",errmsg.ERROR
	}
	return  token ,errmsg.SUCCESS
}

//验证token
func CheckToken(token string) (*MyClaims,int) {
	setToken,_:= jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key  := setToken.Claims.(*MyClaims); setToken.Valid {
		return key ,errmsg.SUCCESS
	} else {
		return nil ,errmsg.ERROR
	}

}
var code int
//jwt中间件
func JwtToke() gin.HandlerFunc  {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader," ",2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.Abort()
			return
		}
		
		key , tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.Abort()
			return
		}

		c.Set("username",key.Username)
		c.Next()
	}
}