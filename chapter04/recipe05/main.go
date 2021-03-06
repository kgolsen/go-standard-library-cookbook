package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2017, 11, 29, 21, 0, 0, 0, time.Local)
	fmt.Printf("Extracting units from time: %v\n", t)

	dayOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()

	fmt.Printf("The %dth day of %v is %v\n",
		dayOfMonth, month, weekDay)

}
