package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2000, 1, 1, 0, 0, 0, 0, l)
	t2 := time.Date(2000, 1, 3, 0, 0, 0, 0, l)
	fmt.Printf("First date: %v\n", t)
	fmt.Printf("Second date: %v\n", t2)

	dur := t2.Sub(t)
	fmt.Printf("Duration between t and t2: %v\n", dur)

	dur = time.Since(t)
	fmt.Printf("Duration between now and t: %v\n", dur)

	dur = time.Until(t)
	fmt.Printf("Duration between t and now: %v\n", dur)

}
