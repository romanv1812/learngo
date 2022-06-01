package main

import "fmt"

func main() {
	// ---------------------------------------------------------
	// EXERCISE: Do Some Calculations
	//
	//  1. Print the sum of 50 and 25
	//  2. Print the difference of 50 and 15.5
	//  3. Print the product of 50 and 0.5
	//  4. Print the quotient of 50 and 0.5
	//  5. Print the remainder of 25 and 3
	//  6. Print the negation of `5 + 2`
	//
	// EXPECTED OUTPUT
	//  75
	//  34.5
	//  25
	//  100
	//  1
	//  -7
	// ---------------------------------------------------------

	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))
}
