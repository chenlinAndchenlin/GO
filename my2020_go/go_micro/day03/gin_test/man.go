package main

import (
	"github.com/gin-gonic/gin"
)

func CallRemote(ctx *gin.Context) {

}
func main() {
	router := gin.Default()
	/*router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("hello chenlin")
		fmt.Println("gin test")
	})*/
	router.GET("/", CallRemote)
	router.Run(":8080")
}
