package errmsg

const (
	SUCCESS = 200
	ERROR = 500

	//code = 1xxx ... 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXISTS   = 1003
	ERROR_TOKEN_NOT_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT = 1008

	// code = 2xxx... 分类模块
	ERROR_CATEGORY_USED = 2001
	ERROR_CATEGORY_NOT_EXISTS = 2002

	// code = 3xxx 文章模块
	ERROR_ARTICLE_NOT_EXISTS = 3001
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR: "FAIL",
	ERROR_USERNAME_USED: "用户名已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXISTS: "用户不存在",
	ERROR_TOKEN_NOT_EXIST: "TOKEN不存在",
	ERROR_TOKEN_RUNTIME: "TOKEN已过期",
	ERROR_TOKEN_WRONG    : "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:"用户没有权限",

	ERROR_CATEGORY_USED: "分类已经存在",
	ERROR_CATEGORY_NOT_EXISTS:"分类不存在",

	ERROR_ARTICLE_NOT_EXISTS:"文章不存在",
}

func GetErrMsg(code int)string  {
	return codeMsg[code]
}
