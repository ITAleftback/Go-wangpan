package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"summerbox/jwt"
)

type User struct {
	gorm.Model
	Username string
	Password string
}
///  model里面放的是数据库层的操作
func (user *User)Register() (message string) {
	DB.Where("username=?",user.Username).Find(&user)
	if user.ID != 0 {
		message="用户名已存在"
		return message
	}
	DB.Create(&user)
	message="注册成功"
	return message
}

func (user *User)Login()(message string)  {
	DB.Where("username=? AND password=?",user.Username,user.Password).Find(&user)
	if user.ID>0{
		message="登录成功"
		//如果登录成功 创建token
		signtoken:=jwt.Create(user.Username,user.ID)
		fmt.Println(signtoken)

		return message
	}else {
		message="密码错误"
		return message
	}
}