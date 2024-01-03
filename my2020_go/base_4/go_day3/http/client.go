package main

import (
	"fmt"
	"net/http"
)

func main() {

	client := http.Client{}

	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("client.Get err:", err)
		return
	}

	/*//beego, gin  ==> web框架
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")
	fmt.Println("header : ", resp.Header)

	fmt.Println("content-type:", ct)
	fmt.Println("date:", date)
	//BWS是Baidu Web Server,是百度开发的一个web服务器 大部分百度的web应用程序使用的是BWS
	fmt.Println("server:", server)
	fmt.Println("*************************************************")

	body := resp.Body
	fmt.Println("body 111:", body)
	//readBodyStr, err := ioutil.ReadAll(body)
	readBodyStr, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("body string:", string(readBodyStr))*/

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url)       //https://www.baidu.com
	fmt.Println("code:", code)     //200
	fmt.Println("status:", status) //OK
}
