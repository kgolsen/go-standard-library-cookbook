package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	io.WriteString(os.Stdout, "Write to STDOUT.\n")

	io.WriteString(os.Stderr, "Write to STDERR.\n")

	// using Writer interface
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// can also use fmt
	fmt.Fprintln(os.Stdout, "From fmt.")
}
