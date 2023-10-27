package main

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var db *sqlx.DB

type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "admin:admin@tcp(127.0.0.1:3306)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	//db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	fmt.Println("连接数据库成功~")
	return nil
}
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := db.Exec(query, args...)
	return err
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	fmt.Println("query:", query)
	fmt.Println("args:", args)
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	fmt.Println("query:", query)
	fmt.Println("args:", args)
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//u1 := User{Name: "xxxx", Age: 18}
	//u2 := User{Name: "xxxxx", Age: 28}
	//u3 := User{Name: "xxxxxx", Age: 38}

	// 方法1
	//users := []*User{&u1, &u2, &u3}
	//err = BatchInsertUsers(users)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	//}

	// 方法2
	//users2 := []interface{}{u1, u2, u3}
	//err = BatchInsertUsers2(users2)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	//}

	// 方法3
	//users3 := []*User{&u1, &u2, &u3}
	//fmt.Println(users3)
	//err = BatchInsertUsers3(users3)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	//}
	ids := []int{1, 2, 10, 14, 5, 9, 15}
	users, err := QueryByIDs(ids)
	if err != nil {
		panic("查询失败")
	}
	for user := range users {
		println(users[user].Name, users[user].Age)

	}
	println("*******************************")
	users, err = QueryAndOrderByIDs(ids)
	if err != nil {
		panic("查询失败")
	}
	for user := range users {
		println(users[user].Name, users[user].Age)

	}
}
