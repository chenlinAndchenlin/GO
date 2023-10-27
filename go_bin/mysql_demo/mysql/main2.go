//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql" //执行init方法 执行驱动注册
//)
//
//type user struct {
//	id   int
//	age  int
//	name string
//}
//
//var db *sql.DB
//
//func initDB() (err error) {
//	dsn := "admin:admin@tcp(127.0.0.1:3306)/go_test"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//
//	db.SetMaxOpenConns(1)
//	db.SetMaxIdleConns(3)
//
//	fmt.Println("连接数据库成功~")
//	return nil
//}
//
//// 预处理插入示例
//func prepareInsertDemo() {
//	sqlStr := "insert into user(name, age) values (?,?)"
//	stmt, err := db.Prepare(sqlStr)
//	if err != nil {
//		fmt.Printf("prepare failed, err:%v\n", err)
//		return
//	}
//	defer stmt.Close()
//	_, err = stmt.Exec("小王子", 18)
//	if err != nil {
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//	_, err = stmt.Exec("沙河娜扎", 18)
//	if err != nil {
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("insert success.")
//}
//
//// 预处理查询示例
//func prepareQueryDemo() {
//	sqlStr := "select id, name, age from user where id > ?"
//	stmt, err := db.Prepare(sqlStr)
//	if err != nil {
//		fmt.Printf("prepare failed, err:%v\n", err)
//		return
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(0)
//	if err != nil {
//		fmt.Printf("query failed, err:%v\n", err)
//		return
//	}
//	defer rows.Close()
//	// 循环读取结果集中的数据
//	for rows.Next() {
//		var u user
//		err := rows.Scan(&u.id, &u.name, &u.age)
//		if err != nil {
//			fmt.Printf("scan failed, err:%v\n", err)
//			return
//		}
//		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
//	}
//}
//
//func main() {
//	err := initDB() // 调用输出化数据库的函数
//	if err != nil {
//		fmt.Printf("init db failed,err:%v\n", err)
//		return
//	}
//
//	defer db.Close() // 必须是在if error后面写，因为如果出现error,db是空的
//
//	prepareQueryDemo()
//	//prepareInsertDemo()
//	//fmt.Println("              ")
//	//prepareQueryDemo()
//}
