package controllers

import (
	"github.com/gin-gonic/gin"
	"summerbox/resps"
	"summerbox/service"
)

//注册===================================================================
func Register(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	message:=service.Register(username,password)
	resps.Ok(c,message)

}


//登录=======================================
func Login(c *gin.Context) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	message:=service.Login(username,password)
	resps.Ok(c,message)
}