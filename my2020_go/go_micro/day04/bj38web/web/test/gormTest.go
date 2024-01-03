package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //"_" 代码不直接使用包, 底层链接要使用!
	"github.com/jinzhu/gorm"
)

func main() {
	// 链接数据库--获取连接池句柄 格式: 用户名:密码@协议(IP:port)/数据库名
	conn, err := gorm.Open("mysql",
		"root:root@tcp(192.168.150.132:3306)/go_test?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	//defer conn.Close()

	GlobalConn = conn

	// 初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	// 最大数
	GlobalConn.DB().SetMaxOpenConns(100)

	// 不要复数表名
	GlobalConn.SingularTable(true)

	// 借助 gorm 创建数据库表.
	fmt.Println(GlobalConn.AutoMigrate(new(Student)).Error)

	// 插入数据
	//InsertData()

	// 查询数据
	//SearchData()

	// 更新数据
	// UpdateData()

	// 删除数据
	//DeleteData()
}
