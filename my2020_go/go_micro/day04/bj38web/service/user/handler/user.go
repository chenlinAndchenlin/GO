package handler

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"time"
	"user/model"
	pb "user/proto"
	"user/utils"
)

type User struct{}

// SendSms(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
// Register(ctx context.Context, in *RegReq, opts ...client.CallOption) (*Response, error)
func (e *User) SendSms(ctx context.Context, req *pb.Request, rsp *pb.Response) error {

	result := model.CheckImgCode(req.Uuid, req.ImgCode)
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
		request.PhoneNumbers = req.Phone
		//随机生成
		rand.Seed(time.Now().UnixNano())
		smsCode := fmt.Sprintf("%06d", rand.Int31n(1000000))
		request.TemplateParam = `{"code":"` + smsCode + `"}`
		fmt.Println("duanxin_str_:", smsCode)
		response, _ := client.SendSms(request)
		if response.IsSuccess() {
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
			//将短信验证码传入redis
			err := model.SaveSmsCode(req.Phone, smsCode)
			if err != nil {
				fmt.Println("存储短信验证码到redis失败：", err)
				rsp.Errno = utils.RECODE_DBERR
				rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			}
		} else {
			rsp.Errno = utils.RECODE_SMSERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		}

	} else {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	return nil
}
func (e *User) Register(ctx context.Context, req *pb.RegReq, rsp *pb.Response) error {
	// 先校验短信验证码,是否正确. redis 中存储短信验证码.
	err := model.CheckSmsCode(req.Mobile, req.SmsCode)
	if err == nil {
		err = model.RegisterUser(req.Mobile, req.Password)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		} else {
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		}
	} else { // 短信验证码错误
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	// 如果校验正确. 注册用户. 将数据写入到 MySQL数据库. 在modelFunc.go里面

	// 短信验证码错误
	return nil
}
