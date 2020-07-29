package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"summerbox/router"
)
func main()  {
	r:=gin.Default()
	router.Registerrouter(r)
	r.Run(":8080")
}
