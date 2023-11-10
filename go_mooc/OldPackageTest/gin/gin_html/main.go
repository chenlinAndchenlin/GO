package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "html_test",
		})
	})
	router.Run(":9090")
}
