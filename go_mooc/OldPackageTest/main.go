package main

import (
	"fmt"
)

type Employment struct {
	Name    string  `json:"name"`
	company Company `json:"company"`
}
type Company struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type PrintResult struct {
	Info string
	Err  error
}

/*
客户端：
 1. 建立连接
    2，将employee对象序列化为json
 3. 发送json
    4.等待服务器发送返回数据
 5. 将服务器返回的数据反序列化威威PrintResrlt的结构体

服务器：
 1. 监听端口
    2，读取二进制数据
 3. 将employee发送的JSON数据反序列化
    4.进行业务逻辑处理
 5. 将处理结果PrintResult序列化为二进制数据
 6. 将数据返回
*/
func PocPrintLn(employment Employment) {

}

// 包管理 异常处理 泛型问题
func main() {
	//fmt.Println(calc.Add(2, 3))
	fmt.Println(Employment{
		Name: "Bobby",
		company: Company{
			Name:    "慕课网",
			Address: "北京市",
		},
	})
}
