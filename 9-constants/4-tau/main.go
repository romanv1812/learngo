package main

import (
	"fmt"
	"math"
)

// ---------------------------------------------------------
// EXERCISE: TAU
//
//  Fix the following program and print the TAU number.
//
// HINT
//  You can use %g verb for printing tau.
//
// EXPECTED OUTPUT
//  tau = 6.283185307179586
// ---------------------------------------------------------

func main() {
	// What's the problem with this code?
	// Why it doesn't work?

	const (
		pi  = math.Pi
		tau = pi * 2
	)

	fmt.Printf("tau = %g\n", tau)
}
