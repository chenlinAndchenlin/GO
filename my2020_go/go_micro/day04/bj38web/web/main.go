package main

import (
	"bj38web/web/controller"
	"bj38web/web/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitRedis()
	_, err := model.InitDb()
	if err != nil {
		fmt.Println("mysqlInit:", err)
		return
	}

	router := gin.Default()

	/*router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目开始了")
	})*/

	router.Static("/home", "view")
	//router.StaticFS("/home", http.Dir("view"))

	//添加路由分组
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
	}

	err = router.Run("192.168.150.1:9090")
	if err != nil {
		println("error:", err)
		return
	}
}
