package main

import (
	"fmt"
	"math"
)

const da = 0.2999999999999999988888878684755769272349
const db = 0.3

func main() {

	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("Strings %s = %s equals: %v\n",
		daStr, dbStr, dbStr == daStr)
	fmt.Printf("Number equals: %v\n", db == da)

	fmt.Printf("Number equals with TOLERANCE: %v\n",
		equals(da, db))
}

const TOLERANCE = 1e-8

func equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < TOLERANCE {
		return true
	}
	return false
}
