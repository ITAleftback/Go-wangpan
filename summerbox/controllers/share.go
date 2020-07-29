package controllers

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)
func ShareQRcode(c *gin.Context)  {
	//这个是拿到想要分享的链接的路径
	path:=c.PostForm("path")
	_ = qrcode.WriteFile(path, qrcode.Medium, 256, "qr.png")
	c.File("qr.png")
}
//MD5加密
func Sharesecret(c *gin.Context){
	//输入想要分享的东西的路径
	path:=c.PostForm("path")
	s:=md5.New()
	s.Write([]byte(path))   // 带加密数据
	signData:=s.Sum(nil)

	c.JSON(200,gin.H{"status":http.StatusOK,"":signData})
}
