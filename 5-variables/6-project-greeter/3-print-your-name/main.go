package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hi %s\n", os.Args[1])
	fmt.Println("How are you?")
}
