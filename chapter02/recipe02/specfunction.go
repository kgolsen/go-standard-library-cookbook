package main

import (
	"fmt"
	"strings"
)

const refString = "Mary*had,a%little_lamb"

func main() {

	// called for each rune in the string
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	words := strings.FieldsFunc(refString, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}
