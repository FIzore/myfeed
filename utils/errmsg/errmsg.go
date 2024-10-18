package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code=1000...用户模块的错误码
	ERROR_USERNAME_USED  = 1001
	ERROR_EMAIL_USED     = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_PASSWORD_WRONG = 1004
	ERROR_EMAIL_EXIST    = 1005
	//code=2000...token模块的错误码
	ERROR_TOKEN_EXIST      = 2001
	ERROR_TOKEN_RUNTIME    = 2002
	ERROR_TOKEN_WRONG      = 2003
	ERROR_TOKEN_TYPE_WRONG = 2004
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_EMAIL_USED:       "邮箱已存在",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_EMAIL_EXIST:      "邮箱已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
