package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

var content = "This is content to check"

func main() {

	checksum := SHA256(content)
	checksum2 := FileSHA256("content.dat")

	fmt.Printf("Checksum 1: %s\n", checksum)
	fmt.Printf("Checksum 2: %s\n", checksum2)
	if checksum == checksum2 {
		fmt.Println("Checksums match.")
	}

}

func SHA256(data string) string {
	h := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", h)
}

func FileSHA256(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	h := sha256.Sum256(b)
	return fmt.Sprintf("%x", h)
}
