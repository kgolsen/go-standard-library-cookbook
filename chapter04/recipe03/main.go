package main

import (
	"fmt"
	"time"
)

func main() {

	// parsing without timezone results in a UTC date
	t, err := time.Parse("2/1/2006", "31/7/2015")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// not the case if you provide a timezone
	t, err = time.Parse("2/1/2006 3:04 PM MST",
		"31/7/2015 1:25 AM PST")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// ParseInLocation. Nifty.
	t, err = time.ParseInLocation("2/1/2006 3:04 PM ",
		"31/7/2015 1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}
