package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create channel for signal to be sent to
	sigChan := make(chan os.Signal, 1)

	// Register signals to catch
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// Create channel wait on signal handling
	exitChan := make(chan int)
	go func() {
		signalReceived := <-sigChan
		switch signalReceived {
		case syscall.SIGHUP:
			fmt.Println("Caught SIGHUP")
			exitChan <- 0

		case syscall.SIGINT:
			fmt.Println("Caught SIGINT")
			exitChan <- 1

		case syscall.SIGTERM:
			fmt.Println("Caught SIGTERM")
			exitChan <- 1

		case syscall.SIGQUIT:
			fmt.Println("Caught SIGQUIT")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
