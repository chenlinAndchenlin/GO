package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("chen?")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.Run()

}
