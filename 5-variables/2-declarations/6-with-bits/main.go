package main

import "fmt"

func main() {
	var (
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64

		f32  float32
		f64  float64
		c64  complex64
		c128 complex128
		off  bool
		s    string
		r    rune
		by   byte
	)
	fmt.Println(i, i8, i16, i32, i64, f32, f64, c64, c128, off, r, by)
	fmt.Printf(" %q\n", s)
}
