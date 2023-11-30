package main

import (
	SUB "12-import/sub" //SUB是我们自己重命名的包名
	//"12-import/sub" //sub是文件名，同时也是包名
	. "12-import/sub" //.代表用户在调用这个包里面的函数时，不需要使用包名.的形式，不见一使用的
	"fmt"
)

func main() {
	//res := sub.Sub(20, 10) //包名.函数去调用
	res := SUB.Sub(20, 10) //包名.函数去调用
	fmt.Println("sub.Sub(20,10) =", res)

	res1 := Sub(30, 20)
	fmt.Println("Sub(30, 20) :", res1)

	//这个无法被调用，是因为函数的首写字母是小写的
	//如果一个包里面的函数想对外提供访问权限，那么一定要首写字母大写,
	// 大写字母开头的函数相当于 public,
	// 小写字母开头的函数相当于 private, 只有相同包名的文件才能使用
	//add.add(10,20)

	fmt.Printf("hello:%s" )
}
