package main

import (
	"fmt"

	"github.com/romanv1812/learngo/learngo/Example/weight"
)

func main() {
	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)
	fmt.Printf("Salt %d,Apples %d, Tuck %d \n", salt, apples, truck)
	salt = Gram(weight.Gram(100))
	fmt.Printf("Type of weight.Gram: %T \n", weight.Gram(1))
}
