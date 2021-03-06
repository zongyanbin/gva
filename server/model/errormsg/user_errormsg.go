package errormsg

const(
	SUCCESS = 2000
	ERROR_SYSTEM_ERROR = 5000
	ERROR_USERNAME_USED = 2001
	ERROR_PASSWORD_WRONG = 2002
	ERROR_USER_NOT_EXIST = 2003
	ERROR_TOKEN_EXIST = 2004
	ERROR_TOKEN_RUNTIME = 2005
	ERROR_TOKEN_WRONG = 2006
	ERROR_TOKEN_TYPE_WRONG = 2007
	ERROR_USER_NO_RIGHT = 2008
)

var codeMsg = map[int]string{
	SUCCESS:					"OK",
	ERROR_USERNAME_USED:		"用户名已存在",
	ERROR_PASSWORD_WRONG:		" 密码错误",
	ERROR_USER_NOT_EXIST:		"用户不存在",
	ERROR_TOKEN_EXIST:			"TOKEN不存在",
	ERROR_TOKEN_RUNTIME:		"TOKEN已过期",
	ERROR_TOKEN_WRONG:			"TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:		"TOKEN格式错误",
	ERROR_USER_NO_RIGHT:		"该用户没权限",
	ERROR_SYSTEM_ERROR:			"系统错误",
}

func GetErrMsg(code int) string  {
	return codeMsg[code]
}