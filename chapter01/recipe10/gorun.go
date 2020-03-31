package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says: " + s.Text())
		}
	}()

	proc.Start()

	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is neat\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()
}
