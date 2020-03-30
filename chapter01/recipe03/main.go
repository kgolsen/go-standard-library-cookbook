package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Custom types need to implement flag.Value to be used in flag.Var().
type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *ArrayValue) Set(str string) error {
	*s = strings.Split(str, ",")
	return nil
}

func main() {
	// flag.Int returns pointer
	retry := flag.Int("retry", -1, "max retry count")

	// Read flags using <Type>Var function. <Type>Var functions take references as the first argument.
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "log prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "comma-delimited array to iterate over")


	// Actually parse options
	flag.Parse()

	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection...")
		logger.Printf("Sending array '%v'\n", arr)
		retryCount++
	}
}
