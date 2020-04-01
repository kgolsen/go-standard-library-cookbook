package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("file.dat")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File name: " + fi.Name())
	fmt.Printf("Is Directory?: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())

}
