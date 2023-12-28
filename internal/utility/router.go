package utility

import "github.com/gin-gonic/gin"

func EnrollRoute(c *gin.RouterGroup) {
	c.GET("/time", Time)
}
