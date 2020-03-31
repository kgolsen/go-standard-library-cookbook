package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Print path to executable
	fmt.Println(exe)

	// Get directory of the executable
	exePath := filepath.Dir(exe)
	fmt.Println("Exe path: ", exePath)

	// Get the real path
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Real path: ", realPath)
}
