package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"myfeed/utils/errmsg"
)

type Accounts struct {
	user_id        uint
	user_name      string
	password       string
	email          string
	xiaohongshu_id string
}

// CheckUser TODO return 使用错误码
func CheckUser(name string) (code int) { //检查用户是否存在
	var account Accounts
	db.Select("user_id").Where("user_name = ?", account.user_name).First(&account)
	if account.user_id > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}
func CheckEmail(email string) (code int) { //检查邮箱是否存在
	var account Accounts
	db.Select("user_id").Where("email = ?", account.email).First(&account)
	if account.user_id > 0 {
		return errmsg.ERROR_EMAIL_EXIST
	}
	return errmsg.SUCCESS
}
func CreateUser(data *Accounts) int {
	data.password = ScryptPassword(data.password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func ChangePassword(id int, data *Accounts) int {
	data.password = ScryptPassword(data.password)
	err := db.Model(&Accounts{}).Where("user_id = ?", id).Update("password", data.password).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
func ScryptPassword(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}
