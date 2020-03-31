package main

import (
	"fmt"
	"strconv"
)

const (
	bin = "10111"
	hex = "1A"
	oct = "012"
	dec = "10"
	flt = 16.123557
)

func main() {

	// convert bin to hex
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("bin2hex(%s): %s\n", bin, v)

	// hex to dec
	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("hex2dec(%s): %s\n", hex, v)

	// oct to hex
	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("oct2hex(%s): %s\n", oct, v)

	// dec to oct
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("dec2oct(%s): %s\n", dec, v)
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
