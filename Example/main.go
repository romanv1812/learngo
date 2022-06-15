package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const (
		feetInMetrs float64 = 0.3048
		feetInYards         = feetInMetrs / 0.9144
	)

	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * feetInMetrs
	yards := math.Round(feet * feetInYards)
	fmt.Printf("%g feet is %g metrs\n", feet, meters)
	fmt.Printf("%g feet is %g yards \n", feet, yards)

}
