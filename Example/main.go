package main

import (
	"fmt"
)

func main() {
	const meters int = 100

	cm := 100
	m := cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)
	cm = 200
	m = cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	const (
		total, number0f int = 5, 1
	)
	fmt.Println(total / number0f)
}
