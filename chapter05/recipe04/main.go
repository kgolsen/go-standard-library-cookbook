package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("### Read as reader ###")
	f, err := os.Open("file.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read with Reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	// Using ReadFile (for smaller files)
	fmt.Println("### ReadFile ###")
	fContent, err := ioutil.ReadFile("file.dat")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))

}
