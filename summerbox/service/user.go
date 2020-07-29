package service

import (

	"summerbox/model"
)

func Register(username string,password string)(message string){
	u:=model.User{
		Username: username,
		Password: password,
	}

	message=u.Register()
	return message

}

func Login(username string,password string)(message string)  {
	u:=model.User{
		Username: username,
		Password: password,
	}

	message=u.Login()

	return message
}
