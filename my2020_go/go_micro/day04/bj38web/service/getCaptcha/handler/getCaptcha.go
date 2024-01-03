package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"getCaptcha/model"
	pb "getCaptcha/proto"
	"github.com/afocus/captcha"
	"image/color"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
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
	// 存储图片验证码到 redis 中
	err := model.SaveImgCode(str_, req.Uuid)
	if err != nil {
		return err
	}
	fmt.Println("tupian_str_:", str_)
	//img, _ := cap.Create(4, captcha.NUM)
	// 将 生成成的图片 序列化.
	imgBuf, _ := json.Marshal(img)

	// 将 imgBuf 使用 参数 rsp 传出
	rsp.Img = imgBuf

	return nil
}
