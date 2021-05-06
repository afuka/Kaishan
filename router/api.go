package router

import (
	"github.com/gin-gonic/gin"
	"time"
)

// InitAPIRouter 初始化 api 的路由
func InitAPIRouter(r *gin.RouterGroup) {
	r.GET("/hello", func(c *gin.Context) {
		time.Sleep(time.Second * 15)
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
}
