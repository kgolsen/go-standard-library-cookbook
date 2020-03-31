package main

import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intStr = "12"
const floatStr = "12.3"

func main() {

	// Decimals
	res, err := strconv.Atoi(intStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	// Hex
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed hex: %d\n", res64)

	// Binary
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed binary: %d\n", resBin)

	// Float
	resFloat, err := strconv.ParseFloat(floatStr, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed float: %.5f\n", resFloat)
}
