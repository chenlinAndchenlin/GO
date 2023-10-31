package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"myBluebell/models"
)

const secret = "liwenzhou.com"

// 吧每一步数据库操作封装为函数
// 等待logic层根据业务需求进行调用
func CheckUserExist(username string) (Is_exists bool, err error) {
	sqlStr := `select count(user_id) from go_test_user where username=?`
	var count int
	err = db.Get(&count, sqlStr, username)
	//查询不到就会有err
	if err != nil {
		Is_exists = false
		return Is_exists, err
	}
	Is_exists = (count > 0)
	//如果查询到了 err就是nil count就是大于0
	return Is_exists, nil
}
func InsertUser(user *models.User) (err error) {
	//执行sql语句入库
	//对密码进行加密，不能铭文直接保存
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into go_test_user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Password, user.Password)
	return

}
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
