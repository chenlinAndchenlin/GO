package main

import "fmt"

func main() {
	p1 := testPtr1()
	fmt.Println("*p1:", *p1)
}

//定义一个函数，返回一个string类型的指针, go语言返回写在参数列表后面
func testPtr1() *string {
	name := "Duke"
	p0 := &name
	fmt.Println("*p0:", *p0)

	city := "深圳"
	ptr := &city
	return ptr
}
