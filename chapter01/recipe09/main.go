package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	proc := exec.Command(cmd, "2")
	proc.Start()
	proc.Wait()

	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Process took: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited successfully?: %t\n", proc.ProcessState.Success())
}
