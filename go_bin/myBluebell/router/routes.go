package router

import (
	"myBluebell/controller"
	"myBluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controller.SingUpHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "注册路由成功...")
	})
	return r
}
