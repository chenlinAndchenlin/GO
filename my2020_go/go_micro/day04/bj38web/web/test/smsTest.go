// // This file is auto-generated, don't edit it. Thanks.
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
//	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
//	util "github.com/alibabacloud-go/tea-utils/v2/service"
//	"github.com/alibabacloud-go/tea/tea"
//	"os"
//	"strings"
//
// )
//
// /**
// * 使用AK&SK初始化账号Client
// * @param accessKeyId
// * @param accessKeySecret
// * @return Client
// * @throws Exception
// */
//
//	func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
//		config := &openapi.Config{
//			// 必填，您的 AccessKey ID
//			AccessKeyId: accessKeyId,
//			// 必填，您的 AccessKey Secret
//			AccessKeySecret: accessKeySecret,
//		}
//		// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
//		config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
//		_result = &dysmsapi20170525.Client{}
//		_result, _err = dysmsapi20170525.NewClient(config)
//		return _result, _err
//	}
//
//	func _main(args []*string) (_err error) {
//		// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
//		// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
//		client, _err := CreateClient(tea.String(os.Getenv("LTAI5t5pcPDVvQKyMjvH5cSo")), tea.String(os.Getenv("PJS5uqdBNy2qhI72mTB9VW326uH421")))
//		if _err != nil {
//			fmt.Println("createClient error: ", _err)
//			return _err
//		}
//
//		sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
//			SignName:      tea.String("陈林果的博客"),
//			TemplateCode:  tea.String("SMS_464060381"),
//			PhoneNumbers:  tea.String("13119471224"),
//			TemplateParam: tea.String("{\"code\":\"1234\"}"),
//		}
//		runtime := &util.RuntimeOptions{}
//		tryErr := func() (_e error) {
//			defer func() {
//				if r := tea.Recover(recover()); r != nil {
//					_e = r
//				}
//			}()
//			// 复制代码运行请自行打印 API 的返回值
//			_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
//			if _err != nil {
//				fmt.Println("SendSmsWithOptions error: ", _err)
//				return _err
//			}
//
//			return nil
//		}()
//
//		if tryErr != nil {
//			var error = &tea.SDKError{}
//			if _t, ok := tryErr.(*tea.SDKError); ok {
//				error = _t
//			} else {
//				error.Message = tea.String(tryErr.Error())
//			}
//			// 错误 message
//			fmt.Println(tea.StringValue(error.Message))
//			// 诊断地址
//			var data interface{}
//			d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
//			d.Decode(&data)
//			if m, ok := data.(map[string]interface{}); ok {
//				recommend, _ := m["Recommend"]
//				fmt.Println(recommend)
//			}
//			_, _err = util.AssertAsString(error.Message)
//			if _err != nil {
//				return _err
//			}
//		}
//		return _err
//	}
//
//	func main() {
//		err := _main(tea.StringSlice(os.Args[1:]))
//		if err != nil {
//			fmt.Println("main error: ", err)
//			panic(err)
//		}
//	}
package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-shanghai", "LTAI5t5pcPDVvQKyMjvH5cSo", "PJS5uqdBNy2qhI72mTB9VW326uH421")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.Domain = "dysmsapi.aliyuncs.com" //域名  ---参考讲义补充!
	//request.PhoneNumbers = "18610382737"
	//request.SignName = "爱家租房网"
	//request.TemplateCode = "SMS_183242785"
	//request.TemplateParam = `{"code":232323}`

	request.SignName = "陈林果的博客"
	request.TemplateCode = "SMS_464060381"
	request.PhoneNumbers = "13119471224"
	request.TemplateParam = "{\"code\":\"1234\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

//func main() {
//	config := sdk.NewConfig()
//
//	// Please ensure that the environment variables ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set.
//	credential := credentials.NewAccessKeyCredential(os.Getenv("LTAI5t5pcPDVvQKyMjvH5cSo"), os.Getenv("PJS5uqdBNy2qhI72mTB9VW326uH421"))
//	/* use STS Token
//	credential := credentials.NewStsTokenCredential(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"), os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN"))
//	*/
//	client, err := dysmsapi.NewClientWithOptions("cn-shanghai", config, credential)
//	if err != nil {
//		fmt.Println("NewClientWithOptions error: ", err)
//		panic(err)
//	}
//
//	request := dysmsapi.CreateSendSmsRequest()
//
//	request.Scheme = "https"
//
//	request.SignName = "陈林果的博客"
//	request.TemplateCode = "SMS_464060381"
//	request.PhoneNumbers = "13119471224"
//	request.TemplateParam = "{\"code\":\"1234\"}"
//
//	response, err := client.SendSms(request)
//	if err != nil {
//		fmt.Println("SendSms error,", err)
//		fmt.Print(err.Error())
//	}
//	fmt.Printf("response is %#v\n", response)
//}
