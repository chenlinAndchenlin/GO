package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Mylogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		context.Set("example", "123456")

		context.Next()

		end := time.Since(t)

		fmt.Printf("耗时：%v\n", end)
		fmt.Println("状态：", context.Writer.Status())

	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			fmt.Println(k, v, token)
		}
	}
}
func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(Mylogger())
	router.Use(TokenRequired())
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong successfully",
		})
	})
	router.Run(":9090")

	//上面等同于router：=gin.Default() 是全局使用

	//authrized := router.Group("/goods")
	//authrized.Use(AuthRequired)

}

//func AuthRequired(context *gin.Context) {
//
//}
