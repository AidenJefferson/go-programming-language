// This program simulates the UNIX echo command
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println("Index:", index, ", Arg:", arg)
	}
}
