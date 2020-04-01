package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// get current permissions
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions: %v\n", fi.Mode())

	// change permissions
	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}
	fi, err = f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions: %v\n", fi.Mode())

}
