package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("There are %d people\n", len(os.Args)-1)
	fmt.Printf("Hello great %s!\n", os.Args[1])
	fmt.Printf("Hello great %s!\n", os.Args[2])
	fmt.Printf("Hello great %s!\n", os.Args[3])
	fmt.Printf("Hello great %s!\n", os.Args[4])
	fmt.Printf("Hello great %s!\n", os.Args[5])
	fmt.Println("Nice to meet you all.")

}
