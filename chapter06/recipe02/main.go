package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	tmpFile, err := ioutil.TempFile("", "foo")
	if err != nil {
		panic(err)
	}
	// caller is responsible for cleaning up
	defer os.Remove(tmpFile.Name())

	fmt.Println(tmpFile.Name())

	// TempDir returns path
	tmpDir, err := ioutil.TempDir("", "foodir")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpDir)
	fmt.Println(tmpDir)

}
