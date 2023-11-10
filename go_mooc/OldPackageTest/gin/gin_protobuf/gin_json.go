package main

import (
	"OldPackageTest/gin/gin_protobuf/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/moreJson", moreJson)
	router.GET("/pureJson", pureJson)
	router.GET("/moreProto", moreProto)

	err := router.Run(":9090")
	if err != nil {
		return
	}
}

func moreProto(context *gin.Context) {
	course := []string{"pythons", "google", "c++"}
	use := &proto.Teacher{
		Name:   "bobby",
		Course: course,
	}
	context.JSON(http.StatusOK, use)
}

func moreJson(c *gin.Context) {
	//var msg struct {
	//	Name    string `json:"user"`
	//	Message string `json:"message"`
	//	Number  int
	//}
	//
	//msg.Message = "this is message test"
	//msg.Number = 20
	//msg.Name = "bobby"
	c.JSON(http.StatusOK, gin.H{
		"html": "<b>hello,chen</b>",
	})

}
func pureJson(c *gin.Context) {
	//var msg struct {
	//	Name    string `json:"user"`
	//	Message string `json:"message"`
	//	Number  int
	//}
	//
	//msg.Message = "this is message test"
	//msg.Number = 20
	//msg.Name = "bobby"
	c.PureJSON(http.StatusOK, gin.H{
		"html": "<b>hello,chen</b>",
	})

}
