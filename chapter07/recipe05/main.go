package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {

	u := &url.URL{
		Scheme:   "http",
		Host:     "localhost",
		Path:     "index.html",
		RawQuery: "id=1&name=John",
		User:     url.UserPassword("admin", "1234"),
	}

	fmt.Printf("Assembled URL:\n%v\n\n", u)

	parsedURL, err := url.Parse(u.String())
	if err != nil {
		panic(err)
	}
	jsonURL, err := json.Marshal(parsedURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed URL:")
	fmt.Println(string(jsonURL))

}
