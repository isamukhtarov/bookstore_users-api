package ping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Function which check if gin working correctly
func Ping(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message": "Pong",
	})
}
