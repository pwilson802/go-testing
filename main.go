package main

import (
	"fmt"
	"os"
	"test-app/messages"
)

func main() {
	// fmt.Printf("Hello, %v\n", os.Args[1])
	fmt.Println(messages.Greet(os.Args[1:]))
}
