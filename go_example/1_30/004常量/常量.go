package main

import (
	"fmt"
	"reflect"
)

func main() {
	const n = 50000

	// fmt.Println("%T", n)
	fmt.Println(reflect.TypeOf(n))
}
