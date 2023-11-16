package api

import (
	"context"
	"fmt"
	"mxshop_api/uer-web/forms"
	"mxshop_api/uer-web/global"
	"mxshop_api/uer-web/global/reponse"
	"mxshop_api/uer-web/proto"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func remove_key(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
func HandlerValidate(c *gin.Context, err error) {

	//如何返回错误信息 翻译和格式化
	//fmt.Println("数据格式不是固定要求的，err:" + err.Error())
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": remove_key(errs.Translate(global.Trans)),
	})
	return
	//c.JSON(http.StatusBadRequest, gin.H{
	//	"message": "请检查数据格式. " + err.Error(),
	//})

	//else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "登录成功",
	//	})
	//}
}
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "404 Not Found:" + e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInsufficientStorage, gin.H{
					"mst": ":用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "其他错误" + e.Message(),
				})
			}
		}
		return
	}
}
func GetUserList(ctx *gin.Context) {
	zap.S().Debug("获取用户列表页")
	ip := global.UserServerConfig.UsersrvInfo.Host
	port := global.UserServerConfig.UsersrvInfo.Port
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
			"message", err.Error())
	}
	//生成grpc接口
	userSrvClient := proto.NewUserClient(conn)
	pn := ctx.DefaultQuery("pn", "0")
	pSize := ctx.DefaultQuery("psize", "10")
	pnInt, _ := strconv.Atoi(pn)
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表】 失败",
			"message", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		//zap.S().Infof("value: %v", value)
		//data := make(map[string]interface{})
		user := reponse.UserResponse{
			Id:       value.Id,
			NickName: value.Nickname,
			//BIrthDay: time.Time(time.Unix(int64(value.BirthDay), 0)),
			BIrthDay: reponse.JsonTiem(time.Unix(int64(value.BirthDay), 0)),
			Mobile:   value.Mobile,
			Gender:   value.Gender,
		}
		//data["id"] = value.Id
		//data["name"] = value.Nickname
		//data["birthday"] = value.BirthDay
		//data["gender"] = value.Gender
		//data["mibile"] = value.Mobile
		//result = append(result, data)
		result = append(result, user)

	}
	ctx.JSON(http.StatusOK, result)
}
func PassWordLogin(c *gin.Context) {
	//表单验证 “generic password”
	passwordLoginForm := forms.PassWordLoginForm{}
	err := c.ShouldBindJSON(&passwordLoginForm)
	if err != nil {
		HandlerValidate(c, err)
		return
	}

	ip := global.UserServerConfig.UsersrvInfo.Host
	port := global.UserServerConfig.UsersrvInfo.Port
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
			"message", err.Error())
	}
	//生成grpc接口
	userSrvClient := proto.NewUserClient(conn)

	//登录逻辑实现
	rsq, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg_mobile": "用户不寻存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg_mobile": "登录失败",
				})
			}
		}
	} else {
		//上面只是查询了用户是否存在，并没有检查密码
		pasRsq, paserr := userSrvClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.Password,
			EncryptedPassword: rsq.Password,
		})
		if paserr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg_password": "登录失败",
			})
		} else if pasRsq.Success {
			c.JSON(http.StatusOK, gin.H{
				"msg_password": "登录成功",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg_password": "密码错误",
			})
		}
	}
}
