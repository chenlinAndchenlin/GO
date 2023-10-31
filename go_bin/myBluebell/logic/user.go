package logic

import (
	"errors"
	"myBluebell/dao/mysql"
	"myBluebell/models"
	"myBluebell/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	var is_exist bool
	if is_exist, err = mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}
	//查询已经存在
	if is_exist {
		return errors.New("用户已存在")
	}
	// 生成id
	userID := snowflake.GenID()
	u := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//密码加密等操作
	//构造一个user实例
	// 保存进数据库
	err = mysql.InsertUser(&u)
	return
}
