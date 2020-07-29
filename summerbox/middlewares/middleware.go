package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"summerbox/jwt"
	"summerbox/resps"
)

func User(c *gin.Context) {
	auth:= c.GetHeader("Authorization")
	fmt.Println(auth)
	if len(auth)<7 {
		resps.Error(c, "token error")
		c.Abort()
		return
	}
	token := auth[7:]
	 err := jwt.CheckToken(token)
	if err != nil {
		resps.Error(c, "token error")
		c.Abort()
		return
	}
	c.Next()
	return
}