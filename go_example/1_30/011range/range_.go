package main

import "fmt"

func main() {
	nums := []int{2, 5, 6}
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}

	println("sum:", sum)

	for i, num := range nums {
		if num == 5 {
			fmt.Println("index:", i)
		}
	}

	keys := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range keys {
		fmt.Println(k, " : ", v)
	}

	for k := range keys {
		fmt.Println("key:", k)
	}
	for i, c := range "AaBb" {
		fmt.Println(i, c)
	}

}
