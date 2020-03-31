package main

import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {

	// common way to print decimal numbers
	fmt.Printf("%d\n", integer)

	// always show sign
	fmt.Printf("%+d\n", integer)

	// print in other base: X:16, o:8, b:2, d:10
	fmt.Printf("%X\n", integer)
	fmt.Printf("%#X\n", integer)

	// pad with leading zeroes
	fmt.Printf("%010d\n", integer)

	// left pad with spaces
	fmt.Printf("% 10d\n", integer)

	// right pad
	fmt.Printf("% -10d\n", integer)

	// float
	fmt.Printf("%f\n", floatNum)

	// float with limited precision
	fmt.Printf("%.5f\n", floatNum)

	// float in scientific notation
	fmt.Printf("%e\n", floatNum)

	// float %e for large exponents, %f otherwise
	fmt.Printf("%g\n", floatNum)
}
