package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Optimal Types
//
//  1. Choose the optimal data types for the given situations.
//  2. Print them all
//  3. Try converting them to lesser data types.
//     For example try converting int64 variable to int32.
//     Then observe the result.
//     Search the web why the result is so?
//
// NOTE
//  This is just an exercise for teaching you different
//  data types. Do not apply it to the real-life.
//
//  As I said in the lectures that, premature optimization
//  is not a good thing.
// ---------------------------------------------------------

func main() {
	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	var latter byte = 'A'
	fmt.Println("an english letter:", latter)
	// a non-english letter (search web for: unicode codepoint)
	var unicode rune = 'Ç'
	fmt.Println("a non-english letter:", unicode)

	// a year in 4 digits like 2040
	var year uint16 = 2040
	fmt.Println("Year: ", year)

	// a month in 2 digits: 1 to 12
	var month uint8 = 8
	fmt.Println("month ", month)

	// the speed of the light
	var speedLight uint64 = 670616629
	fmt.Println("the speed of the light: ", speedLight)

	// angle of a circle
	var angle float32 = 35.8
	fmt.Println("angle of a circle:", angle)
	// PI number: 3.141592653589793
	var pi float64 = 3.141592653589793

	fmt.Println("PI number:", pi)
}
