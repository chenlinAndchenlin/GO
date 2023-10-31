package main

import "fmt"

func main() {
	// s := make([]string, 3)
	// fmt.Println("empty: ", s)

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set: ", s)
	// fmt.Println("get: ", s[2])
	// fmt.Println("length: ", len(s))

	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("append: ", s)
	// // fmt.Println("append: ", d)

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("copy:", c)

	// s[1] = "1"
	// fmt.Println(s)
	// fmt.Println(c)

	// c[1] = "2"
	// fmt.Println(s)
	// fmt.Println(c)

	// l := s[2:5]
	// fmt.Println("sl1:", l)
	// fmt.Println(len(l), cap(l))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
