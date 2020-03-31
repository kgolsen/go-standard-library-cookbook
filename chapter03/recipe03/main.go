package main

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {

	intVal := int(valA)
	fmt.Printf("Rounding by cast to int: %v\n", intVal)

	fRound := math.Round(valA)
	fmt.Printf("Rounding by math.Round: %v\n", fRound)
}
