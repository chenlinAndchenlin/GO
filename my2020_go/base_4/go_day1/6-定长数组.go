package main

import (
	"fmt"
	"time"
)

func main() {

	//1-定义，定义一个具有10个数字的数组
	//c语言定义: int nums[10] ={1,2,3,4}
	//go语言定义：
	// 	nums := [10]int{1,2,3,4} (常用方式)
	//  var nums = [10]int{1,2,3,4}
	// var nums [10]int = [10]int{1,2,3,4}

	nums := [10]int{1, 2, 3, 4}

	//2-遍历，方式一
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i, ", j:", nums[i])
	}

	//方式二： for range   ===> python支持
	//key是数组下标，value是数组的值
	for key, value := range nums {
		//key=0, value=1  => nums[0]
		value += 1
		//value全程只是一个临时变量，不断的被重新赋值，修改它不会修改原始数组
		fmt.Println("key:", key, ", value:", value, "num:", nums[key])
	}

	//在go语言中，如果想忽略一个值，可以使用_
	//如果两个都忽略，那么 就不能使用 := ,而应该直接使用 =
	for _, value := range nums {
		fmt.Println("_忽略key, value:", value)
	}

	//不定长数组定义
	//3-使用make进行创建数组

	go test1()

	for {
		fmt.Println("============> 这是主线程....")
		time.Sleep(1000)
	}
}

func test1() {
	for {
		fmt.Println("这是子线程hello world!")
		time.Sleep(1000)
	}
}
