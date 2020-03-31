package main

import (
	"fmt"
	"regexp"
	"strings"
)

const refString = "Mary had a little lamb"
const refStringTwo = "lamb lamb lamb lamb"

func main() {

	// with strings.Replace
	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)

	// with Replacer
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out = replacer.Replace(refString)
	fmt.Println(out)

	// with regex
	regex := regexp.MustCompile("l[a-z]+")
	out = regex.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}
