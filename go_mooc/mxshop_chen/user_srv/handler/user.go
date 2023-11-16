package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"mxshop_chen/user_srv/global"
	"mxshop_chen/user_srv/model"
	"mxshop_chen/user_srv/proto"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

/*
GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
CheckPassword(context.Context, *PasswordCheckInfo) (*CheckPasswordResponse, error)
*/
type UserServer struct {
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func Model2Rsponse(user model.User) proto.UserInfoResponse {
	//在grpc中message中有默认值 不能随便赋值为nil值 否则容易出错
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Nickname: user.NickName,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected) //之后数据分页

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := Model2Rsponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil

}
func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	//通过mobile查询用户
	var user model.User
	result := global.DB.Where(&model.User{
		Mobile: req.Mobile,
	}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在！请检查信息。")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := Model2Rsponse(user)
	return &userInfoRsp, nil
}
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	//通过ID查询用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在！请检查信息。")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := Model2Rsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//新建用户
	//1、查询用户是否存在

	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已经存在，请直接登录！")
	}
	//没有查询到 则user是直接可以自己使用的
	user.Mobile = req.Mobile
	user.NickName = req.NickName

	//密码加密
	options := &password.Options{
		SaltLen:      10,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode(req.Password, options)
	mypassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println("salt:", salt)
	//fmt.Println("encodedPwd:", encodedPwd)
	//fmt.Println("mypassword:", mypassword)

	//check = password.Verify("generic password", salt, encodedPwd, options)
	//passwordInfo := strings.Split(mypassword, "$")
	//fmt.Println("passwordInfo:", passwordInfo)
	//check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println("check:", check)
	user.Password = mypassword

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	userInfoRsq := Model2Rsponse(user)
	return &userInfoRsq, nil

}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	//个人中心更新用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	birthDay := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = global.DB.Save(user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
func (s *UserServer) CheckPassword(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckPasswordResponse, error) {
	//校验密码
	//数据库存在的是加盐和方法之后的密码：
	//mypassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	options := &password.Options{
		SaltLen:      10,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	//fmt.Println("passwordInfo:", passwordInfo)
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	//fmt.Println("check:", check)
	return &proto.CheckPasswordResponse{Success: check}, nil
}
