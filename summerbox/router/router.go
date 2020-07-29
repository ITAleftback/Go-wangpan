package router

import (
	"github.com/gin-gonic/gin"
	"summerbox/controllers"
	"summerbox/middlewares"
)

func Registerrouter(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.Use(middlewares.User)
	{
		router.POST("/uploadfile", controllers.Uploadfile)
		router.POST("/downloadfile", controllers.Downloadfile)
		//生成二维码分享链接
		router.POST("/shareQRcode", controllers.ShareQRcode)
		//加密分享
		router.POST("/sharesecret", controllers.Sharesecret)
	}
}
