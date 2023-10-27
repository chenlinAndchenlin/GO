//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql" //执行init方法 执行驱动注册
//)
//
//var db *sql.DB
//
//func initDB() (err error) {
//	dsn := "123456:123456@tcp(127.0.0.1:3306)/go_test"
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
////type user struct {
////	id   int
////	age  int
////	name string
////}
//
//// 查询单条数据示例
////
////	func queryRowDemo() {
////		sqlStr := "select id, name, age from user where id=?"
////		//定义结构体 方便保存数据
////		var u user
////		// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
////		// 查询id为1的数据
////		//err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
////		row := db.QueryRow(sqlStr, 1)
////		//扫描并且赋值给变量
////		//err := row.Scan(&u.id, &u.name, &u.age)
////		row = db.QueryRow(sqlStr, 2)
////		//扫描并且赋值给变量
////		err := row.Scan(&u.id, &u.name, &u.age)
////		if err != nil {
////			fmt.Printf("scan failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
////	}
////
//// // 查询多条数据示例
////
////	func queryMultiRowDemo() {
////		sqlStr := "select id, name, age from user where id > ?"
////		rows, err := db.Query(sqlStr, 0)
////		if err != nil {
////			fmt.Printf("query failed, err:%v\n", err)
////			return
////		}
////		// 非常重要：关闭rows释放持有的数据库链接
////		defer rows.Close()
////
////		// 循环读取结果集中的数据
////		for rows.Next() {
////			var u user
////			err := rows.Scan(&u.id, &u.name, &u.age)
////			if err != nil {
////				fmt.Printf("scan failed, err:%v\n", err)
////				return
////			}
////			fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
////		}
////	}
////
//// // 插入数据
////
////	func insertRowDemo() {
////		sqlStr := "insert into user(name, age) values (?,?)"
////		ret, err := db.Exec(sqlStr, "王五", 38)
////		if err != nil {
////			fmt.Printf("insert failed, err:%v\n", err)
////			return
////		}
////		theID, err := ret.LastInsertId() // 新插入数据的id
////		if err != nil {
////			fmt.Printf("get lastinsert ID failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("insert success, the id is %d.\n", theID)
////	}
////
//// 更新数据
////func updateRowDemo() {
////	sqlStr := "update user set age=? where id = ?"
////	ret, err := db.Exec(sqlStr, 339, 3)
////	if err != nil {
////		fmt.Printf("update failed, err:%v\n", err)
////		return
////	}
////	n, err := ret.RowsAffected() // 操作影响的行数
////	if err != nil {
////		fmt.Printf("get RowsAffected failed, err:%v\n", err)
////		return
////	}
////	fmt.Printf("update success, affected rows:%d\n", n)
////}
////
////// 删除数据
////func deleteRowDemo() {
////	sqlStr := "delete from user where id = ?"
////	ret, err := db.Exec(sqlStr, 3)
////	if err != nil {
////		fmt.Printf("delete failed, err:%v\n", err)
////		return
////	}
////	n, err := ret.RowsAffected() // 操作影响的行数
////	if err != nil {
////		fmt.Printf("get RowsAffected failed, err:%v\n", err)
////		return
////	}
////	fmt.Printf("delete success, affected rows:%d\n", n)
////}
//
//func main1() {
//	err := initDB() // 调用输出化数据库的函数
//	if err != nil {
//		fmt.Printf("init db failed,err:%v\n", err)
//		return
//	}
//
//	defer db.Close() // 必须是在if error后面写，因为如果出现error,db是空的
//
//	//queryRowDemo()
//	//queryMultiRowDemo()
//	//insertRowDemo()
//
//	//updateRowDemo()
//	//deleteRowDemo()
//
//}
