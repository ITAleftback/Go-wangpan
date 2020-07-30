package controllers

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"summerbox/resps"
	"summerbox/service"
)

//注册===================================================================
func Register(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username==""||password=="" {
		msg:="您输入的用户名或密码为空"
		resps.Error(c,msg)
		return
	}

	if len(password)<5 {
		msg:="密码必须大于5位"
		resps.Error(c,msg)
		return
	}
	//密码强度必须为字⺟⼤⼩写+数字
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, password); !b || err != nil {
		msg:="密码必须包含字母大小写和数字！"
		resps.Error(c,msg)
		return
	}
	if b, err := regexp.MatchString(a_z, password); !b || err != nil {
		msg:="密码必须包含字母大小写和数字！"
		resps.Error(c,msg)
		return
	}
	if b, err := regexp.MatchString(A_Z, password); !b || err != nil {
		msg:="密码必须包含字母大小写和数字！"
		resps.Error(c,msg)
		return
	}

	message:=service.Register(username,password)
	resps.Ok(c,message)

}


//登录=======================================
func Login(c *gin.Context) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")

	if username==""||password=="" {
		msg:="您输入的用户名或密码为空"
		resps.Error(c,msg)
		return
	}

	message:=service.Login(username,password)
	resps.Ok(c,message)
}