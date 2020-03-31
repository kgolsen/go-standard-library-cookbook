package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	pid := os.Getpid()
	fmt.Printf("Process ID: %d\n", pid)

	proc := exec.Command("ps", "-W")
	out, err := proc.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
