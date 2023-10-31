package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) (d int) {
	d = a + b + c
	return
}
func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
