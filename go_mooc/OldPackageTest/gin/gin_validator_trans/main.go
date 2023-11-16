package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	"github.com/gin-gonic/gin"
)

var trans ut.Translator

func remove_key(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
func InitTrans(locale string) (err error) {
	//修改gin里面的validator；引擎属性
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			err := en_translations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				return err
			}
		case "zh":
			err := zh_translations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				return err
			}
		default:
			err := en_translations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				return err
			}
		}
	}
	return
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=3,max=100"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
type SingUpForm struct {
	Age        uint8  `json:"age" binding:"required,gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3,max=25"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

func main() {
	err := InitTrans("zh")
	if err != nil {
		fmt.Println("翻译初始化错误")
		return
	}
	router := gin.Default()
	//fmt.Println("创建默认的gin")
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginUser Login
		err := c.ShouldBind(&loginUser)
		if err != nil {
			fmt.Println("数据格式不是固定要求的，err:" + err.Error())
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": remove_key(errs.Translate(trans)),
			})
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请检查数据格式. " + err.Error(),
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
			})
		}

	})
	router.POST("/singup", func(c *gin.Context) {
		var singUpform SingUpForm
		err := c.ShouldBind(&singUpform)
		if err != nil {
			fmt.Println("数据格式不是固定要求的，err:" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请检查数据格式. " + err.Error(),
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "注册成功",
			})
		}
	})
	err = router.Run(":9090")
	if err != nil {
		fmt.Println("监听：" + err.Error())
		return
	}
}
