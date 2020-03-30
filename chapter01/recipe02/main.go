package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// Print all arguments
	fmt.Println(args)

	// The first element of args is the called program name
	programName := args[0]
	fmt.Printf("Program called as: %s\n", programName)

	// the rest of the arguments follow...
	otherArgs := args[1:]
	for idx, arg := range otherArgs {
		fmt.Printf("Argument %d: %s\n", idx, arg)
	}
}
