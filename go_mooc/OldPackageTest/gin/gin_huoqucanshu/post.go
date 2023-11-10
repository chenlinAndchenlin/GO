package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", welcome)
	router.POST("/post", formPost)
	router.POST("get_post", get_post)
	err := router.Run(":9090")
	if err != nil {
		return
	}
}

func get_post(context *gin.Context) {
	id := context.Query("ID")
	page := context.DefaultQuery("page", "111")
	name := context.PostForm("name")
	message := context.DefaultPostForm("message", "chenlin de 书 ")
	context.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})

}

func formPost(context *gin.Context) {
	message := context.PostForm("message_chen")
	nick := context.DefaultPostForm("nick", "someone")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(context *gin.Context) {
	firstName := context.DefaultQuery("firstName", "bobby")
	lastName := context.DefaultQuery("lastName", "imooc")

	context.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})

}
