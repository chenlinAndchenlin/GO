package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.150.132:6379")
	if err != nil {
		fmt.Println("redis Dial err:", err)
		return
	}
	defer conn.Close()
	// 2. 操作数据库
	reply, err := conn.Do("set", "itcast", "itheima")
	if err != nil {
		fmt.Println("redis do err:", err)
		return
	}
	// 3. 回复助手类函数. ---- 确定成具体的数据类型
	r, e := redis.String(reply, err)

	fmt.Println(r, e)

}
