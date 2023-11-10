package controller

import (
	"fmt"
	"myBluebell/logic"
	"myBluebell/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SingUpHandler(c *gin.Context) {
	// 1、获取参数 参数校验
	//前后端分离的格式，发过来的数据是json格式的数据
	var p models.ParamSignUp
	//var p :=new(models.ParamSignUp)
	//后面就直接传数据p 不进行取地址
	//ShouldBindJSON功能比较弱，例如只能判断是否值类型是对应的，是否都是JSON，但是比如少一个字段 是检测不出来的
	//前端虽然会进行JSON校验 但是不安全
	//因此需要手动的判断
	if err := c.ShouldBindJSON(&p); err != nil { // 有了binding后 就会进行是否全部参数都有 进行校验 不需要后面的手动校验
		// validator 是有翻译器的 做到了国际化
		zap.L().Error("singup with invalid param", zap.Error(err))
		//判断err是不是validator类型，
		//errs, ok := err.(validator.ValidationErrors)
		if errs, ok := err.(validator.ValidationErrors); !ok {
			c.JSON(http.StatusOK, gin.H{
				//"msg": "请求参数有误",
				"msg": err.Error(), //记录详细的信息
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				//"msg": "请求参数有误",
				"msg": removeTopStruct(errs.Translate(trans)), //记录详细的信息
			})
			return
		}

	}
	//手动的对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("singup with invalid param")
	//	c.JSON(http.StatusOK, gin_hello.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	//上述代码有大量重复的校验使用，因此有现成的包validator参数

	fmt.Println(p)
	//2、业务处理
	err := logic.SignUp(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//3、返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
