package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	Gender string //注意，gender是小写的, 小写字母开头的，在json编码时会忽略掉
}

func main() {

	lily := Student{
		Id:     1,
		Name:   "Lily",
		Age:    20,
		Gender: "女士",
	}
	//虚拟化，struct ==》json
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Println("encodeInfo:", string(encodeInfo))

	//反序列化（解码）： 字符串=》结构体
	var lily2 Student
	if err := json.Unmarshal([]byte(encodeInfo), &lily2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println("name:", lily2.Name)
	fmt.Println("gender:", lily2.Gender)
	fmt.Println("age:", lily2.Age)
	fmt.Println("id:", lily2.Id)
	fmt.Println(lily2)
}
