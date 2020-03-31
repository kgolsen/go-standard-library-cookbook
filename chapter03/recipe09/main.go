package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// complex numbers consist of a real and imaginary component, represented by float64s
	a := complex(2, 3)

	fmt.Printf("Real part: %f\n", real(a))
	fmt.Printf("Imaginary part: %f\n", imag(a))

	// all basic operations are supported
	b := complex(6, 4)
	fmt.Printf("Difference: %v\n", a-b)
	fmt.Printf("Sum: %v\n", a+b)
	fmt.Printf("Product: %v\n", a*b)
	fmt.Printf("Quotient: %v\n\n", a/b)

	fmt.Printf("Conjugate(a): %v\n", cmplx.Conj(a))
	fmt.Printf("Cos(b): %v\n", cmplx.Cos(b))

}
