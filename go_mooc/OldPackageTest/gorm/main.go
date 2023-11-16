package main

//
//import (
//	"fmt"
//	"log"
//	"os"
//	"time"
//
//	"gorm.io/gorm/logger"
//
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type User struct {
//	UserID uint   `gorm:"primarykey"`
//	Name   string `gorm:"column:user_name;type:varchar(50);index:idx_user_name"`
//}
//
//func main() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//	dsn := "root:root@tcp(192.168.150.132:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold:             time.Second, // Slow SQL threshold
//			LogLevel:                  logger.Info, // Log level
//			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
//			ParameterizedQueries:      true,        // Don't include params in the SQL log
//			Colorful:                  true,        // Disable color
//		},
//	)
//
//	// Globally mode
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//	})
//	if err != nil {
//		fmt.Println("数据库连接错误")
//		panic(err)
//	}
//	//设置全局的logger，这个logger在我们执行每个sql中，会打印每一行sql
//	//sql才是最重要的 api背后的sql
//
//	//定义一个表结构，将表结构映射为表 migrations
//	//迁移 schema
//	err = db.AutoMigrate(&User{}) //此处应该有sql语句进行执行
//	if err != nil {
//		fmt.Println("初始化建表错误！！")
//		return
//	}
//	//db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})
//
//	// Read
//	//var product Product
//	//db.First(&product, 1)                 // 根据整型主键查找
//	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
//	//
//	//// Update - 将 product 的 price 更新为 200
//	//db.Model(&product).Update("Price", 200)
//	//var product Product
//	//db.First(&product, 1) // 根据整型主键查找
//	//db.First(&product, "code = ?", "D42")
//	// Update - 更新多个字段
//	//db.Model(&product).Updates(Product{Price: 800, Code: sql.NullString{"", true}}) // 仅更新非零值字段
//	//db.Model(&product).Updates(map[string]interface{}{"Price": 500, "Code": "G42"})
//
//	// Delete - 删除 product
//	//db.Delete(&product, 1)
//	//db.Create(&Product{Code: "L1212", Price: 1000})
//
//}
