package main

import (
	"fmt"
	"os"
)

func main() {
	cmds := os.Args
	for key, cmd := range cmds {
		fmt.Println(key, cmd)
	}

}
