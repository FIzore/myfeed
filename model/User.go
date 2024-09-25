package model

import (
// "gorm.io/gorm"
)

type User struct {
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password "validate:"required,min=6,max=100" label:"密码"`
	Email    string `gorm:"type:varchar(100);not null" json:"email" validate:"required,email" label:"邮箱"`
	//Role     int
}

// CheckUser TODO return 使用错误码
//func CheckUser(name string) string { //检查用户是否存在
//	var user User
//	db.Select("id").Where("username = ?", name).First(&user)
//	if user.ID > 0 {
//		return "用户名已存在"
//	}
//	return "用户不存在"
//}

//func CreateUser() string {
//
//}
//func ResetPassword() string {
//
//}
//func SendEmail() string {
//
//}
