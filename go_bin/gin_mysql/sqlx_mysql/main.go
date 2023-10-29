package main

//
//import (
//	"errors"
//	//"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql" //执行init方法 执行驱动注册
//	"github.com/jmoiron/sqlx"
//)
//
//type user struct {
//	ID   int    `db:"id"`
//	Age  int    `db:"age"`
//	Name string `db:"name"`
//}
//
//// var db *sql.DB
//var db *sqlx.DB
//
//func initDB() (err error) {
//	dsn := "admin:admin@tcp(127.0.0.1:3306)/go_test"
//	db, err = sqlx.Connect("mysql", dsn)
//	//db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		fmt.Printf("connect DB failed, err:%v\n", err)
//		return
//	}
//
//	db.SetMaxOpenConns(20)
//	db.SetMaxIdleConns(10)
//
//	fmt.Println("连接数据库成功~")
//	return nil
//}
//
//// // 查询单条数据示例
////
////	func queryRowDemo() {
////		sqlStr := "select id, name, age from user where id=?"
////		var u user
////		// 使用get进行查询 可以直接赋值给u结构体 这里要指针地址进行传递
////		err := db.Get(&u, sqlStr, 1)
////		if err != nil {
////			fmt.Printf("get failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
////	}
////
//// // 查询多条数据示例
////
////	func queryMultiRowDemo() {
////		sqlStr := "select id, name, age from user where id > ?"
////		var users []user
////		err := db.Select(&users, sqlStr, 0)
////		if err != nil {
////			fmt.Printf("query failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("users:%#v\n", users)
////	}
////
//// // 插入数据
////
////	func insertRowDemo() {
////		sqlStr := "insert into user(name, age) values (?,?)"
////		ret, err := db.Exec(sqlStr, "沙河小王子", 189)
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
////		n, err := ret.RowsAffected() // 操作影响的行数
////		if err != nil {
////			fmt.Printf("get RowsAffected failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("get RowsAffected is %d.\n", n)
////	}
////
//// // 更新数据
////
////	func updateRowDemo() {
////		sqlStr := "update user set age=? where id = ?"
////		ret, err := db.Exec(sqlStr, 39, 6)
////		if err != nil {
////			fmt.Printf("update failed, err:%v\n", err)
////			return
////		}
////		n, err := ret.RowsAffected() // 操作影响的行数
////		if err != nil {
////			fmt.Printf("get RowsAffected failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("update success, affected rows:%d\n", n)
////	}
////
//// // 删除数据
////
////	func deleteRowDemo() {
////		sqlStr := "delete from user where id = ?"
////		ret, err := db.Exec(sqlStr, 6)
////		if err != nil {
////			fmt.Printf("delete failed, err:%v\n", err)
////			return
////		}
////		n, err := ret.RowsAffected() // 操作影响的行数
////		if err != nil {
////			fmt.Printf("get RowsAffected failed, err:%v\n", err)
////			return
////		}
////		fmt.Printf("delete success, affected rows:%d\n", n)
////	}
////
////	func insertUserDemo() (err error) {
////		sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
////		_, err = db.NamedExec(sqlStr,
////			map[string]interface{}{
////				"name": "七qi米",
////				"age":  28,
////			})
////		return
////	}
////
////	func namedQuery() {
////		sqlStr := "SELECT * FROM user WHERE name=:name"
////		// 使用map做命名查询
////		rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七qi米"})
////		if err != nil {
////			fmt.Printf("db.NamedQuery failed, err:%v\n", err)
////			return
////		}
////		defer rows.Close()
////		for rows.Next() {
////			var u user
////			err := rows.StructScan(&u)
////			if err != nil {
////				fmt.Printf("scan failed, err:%v\n", err)
////				continue
////			}
////			fmt.Printf("user:%#v\n", u)
////		}
////
////		u := user{
////			Name: "七11米",
////		}
////		// 使用结构体命名查询，根据结构体字段的 db tag进行映射
////		rows, err = db.NamedQuery(sqlStr, u)
////		if err != nil {
////			fmt.Printf("db.NamedQuery failed, err:%v\n", err)
////			return
////		}
////		defer rows.Close()
////		for rows.Next() {
////			var u user
////			err := rows.StructScan(&u)
////			if err != nil {
////				fmt.Printf("scan failed, err:%v\n", err)
////				continue
////			}
////			fmt.Printf("user:%#v\n", u)
////		}
////	}
//func transactionDemo2() (err error) {
//	tx, err := db.Beginx() // 开启事务
//	if err != nil {
//		fmt.Printf("begin trans failed, err:%v\n", err)
//		return err
//	}
//	defer func() {
//		if p := recover(); p != nil {
//			tx.Rollback()
//			panic(p) // re-throw panic after Rollback
//		} else if err != nil {
//			fmt.Println("rollback")
//			tx.Rollback() // err is non-nil; don't change it
//		} else {
//			err = tx.Commit() // err is nil; if Commit returns error update err
//			fmt.Println("commit")
//		}
//	}()
//
//	sqlStr1 := "Update user set age=220 where id=?"
//
//	rs, err := tx.Exec(sqlStr1, 1)
//	if err != nil {
//		return err
//	}
//	n, err := rs.RowsAffected()
//	if err != nil {
//		return err
//	}
//	if n != 1 {
//		return errors.New("exec sqlStr1 failed")
//	}
//	sqlStr2 := "Update user set age=450 where id=?"
//	rs, err = tx.Exec(sqlStr2, 5)
//	if err != nil {
//		return err
//	}
//	n, err = rs.RowsAffected()
//	if err != nil {
//		return err
//	}
//	if n != 1 {
//		return errors.New("exec sqlStr1 failed")
//	}
//	return err
//}
//
//func main() {
//	if err := initDB(); err != nil {
//		fmt.Printf("init DB failed, err:%v\n", err)
//		return
//	}
//
//	//queryRowDemo()
//	//queryMultiRowDemo()
//	//insertRowDemo()
//	//insertUserDemo()
//	//namedQuery()
//	//namedQuery()
//
//	transactionDemo2()
//
//}
