package mysql

import (
	//"database/sql"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql" //执行init方法 执行驱动注册
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// var db *sql.DB
var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	db, err = sqlx.Connect("mysql", dsn)
	//db, err = sql.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.maxopenconns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.maxidleconns"))

	fmt.Println("连接数据库成功~")
	return nil
}
func Close() {
	_ = db.Close()
}