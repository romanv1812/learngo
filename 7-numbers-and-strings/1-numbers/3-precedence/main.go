package main

//dsa
import "fmt"

func main() {
	// This expression should print 20
	fmt.Println(10 + 5 - (5 - 10))

	// This expression should print -16
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// This expression should print -25
	fmt.Println(5 + 10*(2-5))

	// This expression should print 0.5
	fmt.Println(0.5 * (2 - 1))

	// This expression should print 24
	fmt.Println((3+1)/2*10 + 4)

	// This expression should print 15
	fmt.Println((10 / 2) * (10 % 7))

	// This expression should print 40
	// Note that you may need to use floats to solve this
	fmt.Println(100 / (5. / 2))
}
