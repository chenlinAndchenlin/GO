package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	//优雅退出 就是指退出程序后应该做的后续的处理

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ping",
		})
	})

	//router.Run(":9090")

	go func() {
		router.Run(":9090")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//处理后续的逻辑

	fmt.Println("关闭server中")
	fmt.Println("注销服务中")
}
