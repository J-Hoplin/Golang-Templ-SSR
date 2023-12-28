package todo

import "github.com/gin-gonic/gin"

func EnrollRoute(group *gin.RouterGroup) {
	group.GET("/", ListTODO)
	group.GET("/:tid")
	group.POST("/")
	group.PATCH("/:tid")
	group.DELETE("/:tid", DeleteTODO)
}
