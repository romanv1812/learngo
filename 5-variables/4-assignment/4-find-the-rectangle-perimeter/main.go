package main

import "fmt"

func main() {

	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Printf("Perimeter = %d\n", perimeter)

}
