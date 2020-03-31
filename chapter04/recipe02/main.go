package main

import (
	"fmt"
	"time"
)

func main() {

	tTime := time.Date(1978, time.May, 19, 12, 17, 17, 1010101010, time.Local)

	// formatting can be done with the reference value:
	// Jan 2 15:04:05 2006 MST
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))
	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))

	// can also use predefined formats
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))

	// formatting supports space padding for days
	fmt.Printf("tTime (padded) is: %s\n", tTime.Format("2006/1/_2"))

	// can also do zero padding
	fmt.Printf("tTime (0-padded) is: %s\n", tTime.Format("2006/01/02"))

	// format fractional time component with leading zeroes
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.000"))

	// format fractional time component without leading zeroes
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))

	// append formatted time to given buffer
	fmt.Println(string(tTime.AppendFormat([]byte("The time is now: "), "03:04PM")))
}
