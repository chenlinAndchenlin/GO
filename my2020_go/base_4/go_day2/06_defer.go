package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "1_switch.go"
	readFile(fileName)
}

func readFile(fileName string) {

	open, err := os.Open(fileName)
	if err != nil {
		fmt.Println("os.Open(\"01-switch.go\") ==> 打开文件失败, err:", err)

		return
	}
	defer func(a int) {
		fmt.Println("准备关闭switch文件, code:", a)
		_ = open.Close()
	}(100)

	defer fmt.Println("0000000")
	defer fmt.Println("1111111")
	defer fmt.Println("2222222")

	buff := make([]byte, 1024)
	n, err := open.Read(buff)
	if err != nil {
		return
	}
	fmt.Println("读取长度：", n)
	fmt.Println("duquneir：", string(buff))

}
