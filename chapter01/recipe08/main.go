package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	proc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	proc.Stdout = out
	err := proc.Run()
	if err != nil {
		panic(err)
	}

	if proc.ProcessState.Success() {
		fmt.Println("ls ran successfully with output:")
		fmt.Println(out.String())
	}
}
