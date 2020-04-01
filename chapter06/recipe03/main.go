package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go is neat\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Sure! Go is pretty neat.\n"))
	if err != nil {
		panic(err)
	}

}
