package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there are any command-line arguments
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s!\n", os.Args[1])
	} else {
		fmt.Println("Hello, World!")
	}
}