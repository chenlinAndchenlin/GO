package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"action" binding:"required"`
}

func main() {
	router := gin.Default()
	goodsGroup := router.Group("goods")
	{
		//goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/:id/:action", goodsDetail)
		goodsGroup.POST("/add", createGoods)
	}

	//v1 := router.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login", LoginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//	//router.GET("/goods/list", goodsList)
	//	//router.GET("/goods/1", goodDetails)
	//	//router.GET("goods/add", goodAdd)
	//}

	err := router.Run(":9091")
	if err != nil {
		return
	}
}

func createGoods(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "createGoods"})

}

func goodsDetail(context *gin.Context) {
	var person Person
	err := context.ShouldBindUri(&person)
	if err != nil {
		println("路由数据错误")
		context.Status(404)
		return
	}
	//action := context.Param("action")
	//id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{"message": "goodsDetail", "id": person.ID, "action": person.Name})

}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "goodsList"})
}
