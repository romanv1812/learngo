package main

import "fmt"

func main() {

	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Printf("Red: %s Blue: %s\n", red, blue)
}
