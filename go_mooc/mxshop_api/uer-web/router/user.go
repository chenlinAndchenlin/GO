package router

import (
	"mxshop_api/uer-web/api"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		zap.S().Info("配置用户相关的url:")
		UserRouter.GET("list", api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
	}

}
