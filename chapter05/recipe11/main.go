package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var buf bytes.Buffer

	// compress
	zipW := zip.NewWriter(&buf)
	f, err := zipW.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("This is some file content"))
	if err != nil {
		panic(err)
	}
	err = zipW.Close()
	if err != nil {
		panic(err)
	}

	// write to file
	err = ioutil.WriteFile("data.zip", buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// decompress
	zipR, err := zip.OpenReader("data.zip")
	if err != nil {
		panic(err)
	}

	for _, file := range zipR.File {
		fmt.Println("File " + file.Name + " contains:")
		r, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		if _, err = io.Copy(os.Stdout, r); err != nil {
			panic(err)
		}
		if err = r.Close(); err != nil {
			panic(err)
		}
		fmt.Println()
	}

}
