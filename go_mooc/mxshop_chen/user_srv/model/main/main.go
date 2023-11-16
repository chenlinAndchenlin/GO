package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func genMD5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}
func main() {
	//// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:root@tcp(192.168.150.132:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		//ParameterizedQueries:      true,          // Don't include params in the SQL log
	//		Colorful: true, // Disable color
	//	},
	//)
	//
	//// Globally mode
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	fmt.Println("数据库连接错误")
	//	panic(err)
	//}
	////设置全局的logger，这个logger在我们执行每个sql中，会打印每一行sql
	////sql才是最重要的 api背后的sql
	//
	////定义一个表结构，将表结构映射为表 migrations
	////迁移 schema
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	return
	//}
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println("salt:", salt)
	//fmt.Println("encodedPwd:", encodedPwd)
	//
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println("check:", check)
	//
	//fmt.Println("*******************************")
	//options := &password.Options{
	//	SaltLen:      10,
	//	Iterations:   100,
	//	KeyLen:       32,
	//	HashFunction: sha512.New,
	//}
	//salt, encodedPwd := password.Encode("generic password", options)
	//mypassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println("salt:", salt)
	//fmt.Println("encodedPwd:", encodedPwd)
	//fmt.Println("mypassword:", mypassword)
	//
	////check = password.Verify("generic password", salt, encodedPwd, options)
	//passwordInfo := strings.Split(mypassword, "$")
	//fmt.Println("passwordInfo:", passwordInfo)
	//check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println("check:", check)

	//options := &password.Options{
	//	SaltLen:      10,
	//	Iterations:   100,
	//	KeyLen:       32,
	//	HashFunction: sha512.New,
	//}
	//salt, encodedPwd := password.Encode("generic password", options)
	//mypassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	////fmt.Println("salt:", salt)
	////fmt.Println("encodedPwd:", encodedPwd)
	//fmt.Println("mypassword:", mypassword)
	//for i := 0; i < 10; i++ {
	//	user := model.User{
	//		NickName: fmt.Sprintf("bobby%d", i),
	//		Mobile:   fmt.Sprintf("1878222222%d", i),
	//		Password: mypassword,
	//	}
	//	global.DB.Save(&user)
	//}
}
