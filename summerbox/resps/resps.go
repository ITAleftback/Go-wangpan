package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context,message string) {
	c.JSON(200, gin.H{"status":http.StatusOK, "message": message})
}


func Error(c *gin.Context, msg string) {
	c.JSON(500, gin.H{"status":http.StatusInternalServerError, "message": msg})
}
