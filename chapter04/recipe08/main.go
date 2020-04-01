package main

import (
	"fmt"
	"time"
)

func main() {

	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2000, 1, 1, 0, 0, 0, 0, eur)
	fmt.Printf("Original date: %v\n", t)

	phx, err := time.LoadLocation("America/Phoenix")
	if err != nil {
		panic(err)
	}

	// Okay, this is pretty cool
	t2 := t.In(phx)
	fmt.Printf("In Phoenix: %v\n", t2)

}
