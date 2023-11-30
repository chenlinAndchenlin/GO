package main

import "fmt"

func main() {
	names := [7]string{"北京", "上海", "广州", "深圳", "洛阳", "南京", "秦皇岛"}

	//想基于names创建一个新的数组
	names1 := [3]string{}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]

	//切片可以基于一个数组，灵活的创建新的数组
	names2 := names[0:3] //左闭右开
	fmt.Println("names2:", names2)

	names2[2] = "Hello"
	fmt.Println("修改names[2]之后, names2:", names2)
	fmt.Println("修改names[2]之后, names:", names)

	//1. 如果从0元素开始截取，那么冒号左边的数字可以省略
	names3 := names[:5]
	fmt.Println("name3:", names3)

	//2. 如果截取到数组最后一个元素，那么冒号右边的数字可以省略
	names4 := names[5:]
	fmt.Println("name4:", names4)

	//3. 如果想从左至右全部使用，那么冒号左右两边的数字都可以省略
	names5 := names[:]
	fmt.Println("name5:", names5)

	//4. 也可以基于一个字符串进行切片截取 : 取字符串的字串 helloworld
	sub1 := "helloworld"[5:7]
	fmt.Println("sub1:", sub1) //'wo'

	//5. 可以在创建空切片的时候，明确指定切片的容量，这样可以提高运行效率
	//创建一个容量是20，当前长度是0的string类型切片
	//make的时候，初始的值都是对应类型的零值 : bool ==> false, 数字==> 0, 字符串 ==> 空
	str2 := make([]string, 10, 20) //第三个参数不是必须的，如果没填写，则默认与长度相同
	fmt.Println("str2:", &str2[0])

	fmt.Println("str2: len:", len(str2), ", cap:", cap(str2))
	str2[0] = "hello"
	str2[1] = "world"

	//6.如果想让切片完全独立于原始数组，可以使用copy()函数来完成
	namesCopy := make([]string, len(names))
	//func copy(dst, src []Type) int
	//names是一个数组，copy函数介收的参数类型是切片，所以需要使用[:]将数组变成切片
	copy(namesCopy, names[:])
	namesCopy[0] = "香港"
	fmt.Println("namesCopy:", namesCopy)
	fmt.Println("naemes本身:", names)

}
