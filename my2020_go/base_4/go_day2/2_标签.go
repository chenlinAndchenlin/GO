package main

import "fmt"

func main() {
	//LABEL121:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL121
				//continue LABEL121
				//break LABEL121
				break
			}

			fmt.Println("i:", i, ",j:", j)
		}
	}

	fmt.Println("over!")

}
