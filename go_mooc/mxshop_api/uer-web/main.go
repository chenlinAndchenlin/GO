package main

import (
	"fmt"
	"mxshop_api/uer-web/global"
	"mxshop_api/uer-web/initialize"
	"mxshop_api/uer-web/validator"

	ut "github.com/go-playground/universal-translator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

func main() {
	//1 初始化logger、配置文件、翻译器、自定义验证器
	initialize.InitLogger()
	initialize.InitConfig()
	err := initialize.InitTrans("zh")
	if err != nil {
		fmt.Println("翻译初始化错误")
		return
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myValidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}手机号码不符合规则", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
		zap.S().Infof("mobile注册器完成")
	}

	//2 初始化router、
	Router := initialize.Routers()

	//zap.S().Info("启动服务器！")
	port := global.UserServerConfig.Port
	//zap.S().Infof("端口：%d", port)
	zap.S().Debugf("启动服务器！端口：%d", port)
	err = Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
