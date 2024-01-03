package controller

import (
	getCaptcha "bj38web/web/proto/getCaptcha"
	userMicro "bj38web/web/proto/user"
	"bj38web/web/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"image/png"
	"net/http"
)

// 获取 session 信息.
func GetSession(ctx *gin.Context) {
	// 初始化错误返回的 map
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

/*
	func GetImageCd(ctx *gin.Context) {
		// 获取图片验证码 uuid
		uuid := ctx.Param("uuid")
		// 初始化对象
		cap := captcha.New()

		// 设置字体
		cap.SetFont("./conf/comic.ttf")

		// 设置验证码大小
		cap.SetSize(128, 64)

		// 设置干扰强度
		cap.SetDisturbance(captcha.MEDIUM)

		// 设置前景色
		cap.SetFrontColor(color.RGBA{0, 0, 0, 255})

		// 设置背景色
		cap.SetBkgColor(color.RGBA{100, 0, 255, 255}, color.RGBA{255, 0, 127, 255}, color.RGBA{255, 255, 10, 255})

		// 生成字体 -- 将图片验证码, 展示到页面中.

		img, str_ := cap.Create(4, captcha.NUM)
		png.Encode(ctx.Writer, img)

		fmt.Println("uuid = ", uuid)
		fmt.Println("str = ", str_)
	}
*/
func GetImageCd(ctx *gin.Context) {
	// 获取图片验证码 uuid
	uuid := ctx.Param("uuid")
	// 指定 consul 服务发现
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})
	//consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)
	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getcaptcha", consulService.Client())
	// 调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.CallRequest{Uuid: uuid})
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("短信验证码未找到远程服务...")
		return
	}

	// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 将图片写出到浏览器.
	png.Encode(ctx.Writer, img)

	fmt.Println("uuid = ", uuid)
	//fmt.Println("str = ", str_)
}

// 获取短信验证码
/*func GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	phone := ctx.Param("phone")
	// 拆分 GET 请求中 的 URL === 格式: 资源路径?k=v&k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")
	resp := make(map[string]string)
	result := model.CheckImgCode(uuid, imgCode)
	if result {
		client, _ := dysmsapi.NewClientWithAccessKey("cn-shanghai", "LTAI5t5pcPDVvQKyMjvH5cSo", "PJS5uqdBNy2qhI72mTB9VW326uH421")

		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https"

		request.Domain = "dysmsapi.aliyuncs.com" //域名  ---参考讲义补充!
		//request.PhoneNumbers = "18610382737"
		//request.SignName = "爱家租房网"
		//request.TemplateCode = "SMS_183242785"
		//request.TemplateParam = `{"code":232323}`

		request.SignName = "陈林果的博客"
		request.TemplateCode = "SMS_464060381"
		request.PhoneNumbers = phone
		//随机生成
		rand.Seed(time.Now().UnixNano())
		smsCode := fmt.Sprintf("%06d", rand.Int31n(1000000))
		request.TemplateParam = `{"code":"` + smsCode + `"}`
		fmt.Println("duanxin_str_:", smsCode)
		response, _ := client.SendSms(request)
		if response.IsSuccess() {
			resp["errno"] = utils.RECODE_OK
			resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
			//将短信验证码传入redis
			err := model.SaveSmsCode(phone, smsCode)
			if err != nil {
				fmt.Println("存储短信验证码到redis失败：", err)
				resp["errno"] = utils.RECODE_DBERR
				resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
			}
		} else {
			resp["errno"] = utils.RECODE_SMSERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_SMSERR)
		}

	} else {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	}
	//发送成功或者失败信息
	ctx.JSON(http.StatusOK, resp)

}*/

func GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	phone := ctx.Param("phone")
	// 拆分 GET 请求中 的 URL === 格式: 资源路径?k=v&k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	// 指定 consul 服务发现
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.150.132:8500",
		}
	})
	//consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)
	// 初始化客户端
	microClient := userMicro.NewUserService("user", consulService.Client())
	// 调用远程函数
	resp, err := microClient.SendSms(context.TODO(), &userMicro.Request{Phone: phone, ImgCode: imgCode, Uuid: uuid})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		return
	}
	// 发送校验结果 给 浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 发送注册信息
func PostRet(ctx *gin.Context) {
	/*	mobile := ctx.PostForm("mobile")
		pwd := ctx.PostForm("password")
		sms_code := ctx.PostForm("sms_code")

		fmt.Println("m = ", mobile, "pwd = ", pwd, "sms_code = ",sms_code)*/
	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}

	ctx.Bind(&regData)

	//fmt.Println("获取到的数据为:", regData)
	// 初始化consul
	microService := utils.InitMicro()
	microClient := userMicro.NewUserService("user", microService.Client())
	// 调用远程函数
	resp, err := microClient.Register(context.TODO(), &userMicro.RegReq{
		Mobile:   regData.Mobile,
		SmsCode:  regData.SmsCode,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("注册用户, 找不到远程服务!", err)
		return
	}
	// 写给浏览器
	ctx.JSON(http.StatusOK, resp)
}
