//每个go程序，都必须有一个 package main
//每个go程序，都是.go结尾的
//go程序中没有.h，没有.o, 只有.go
//一个package，包名，相当于命名空间
//std::cout
package main

//这是导入一个标准包fmt，format，一般用于格式化输出
import (
	"fmt"
	"time"
)

//主函数，所有的函数必须使用func 开头
//一个函数的返回值，不会放在func前，而是放在参数后面
//函数左花括号必须与函数名同行，不能写到下一行

func main() {

	//go语言语句不需要使用分号结尾
	fmt.Println("hello world");
	time.Now()
}
