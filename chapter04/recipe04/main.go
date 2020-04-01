package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Unix(0, 0)
	fmt.Println(t)

	epoch := t.Unix()
	fmt.Println(epoch)

	epochNow := time.Now().Unix()
	fmt.Println(epochNow)

	epochNano := time.Now().UnixNano()
	fmt.Println(epochNano)

}
