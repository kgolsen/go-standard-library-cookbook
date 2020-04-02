package main

import (
	"fmt"
	"net/http"
)

func main() {

	header := http.Header{}

	// can use Header as a slice
	header.Set("Auth-X", "1234567890a")
	header.Add("Auth-X", "abcdefghijk")
	fmt.Println(header)

	hdrSlice := header["Auth-X"]
	fmt.Println(hdrSlice)

	// just get the first value
	hdrFirst := header.Get("Auth-X")
	fmt.Println(hdrFirst)

	// using Set again overwrites the existing value(s)
	header.Set("Auth-X", "a0987654321")
	fmt.Println(header)

	// remove the item
	header.Del("Auth-X")
	fmt.Println(header)

}
