package model

type Accounts struct {
	user_id        uint
	user_name      string
	password       string
	email          string
	xiaohongshu_id string
}

// CheckUser TODO return 使用错误码
func CheckUser(name string) string { //检查用户是否存在
	var account Accounts
	db.Select("user_id").Where("user_name = ?", account.user_name).First(&account)
	if account.user_id > 0 {
		return "用户名已存在"
	}
	return "用户不存在"
}

//func CreateUser() string {
//
//}
//func ResetPassword() string {
//
//}
//func SendEmail() string {
//
//}
//todo 加密密码
