package utility

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Time(c *gin.Context) {
	now := time.Now()
	timeString := fmt.Sprintf("%d년 %02d월 %02d일 %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	c.String(http.StatusOK, timeString)
}
