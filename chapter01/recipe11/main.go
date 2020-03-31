package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func main() {

	// Open file as a log file to write to
	var err error
	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log",
		time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// use closeChan to detect program termination
	closeChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-closeChan:
				log.Println("Gorouting closing")
				return
			default:
				log.Println("Writing to log")
				io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now().String()))
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT)

	// blocking read from sigChan, where Notify sends signals
	<-sigChan

	// after signal is received, all code after the read from the channel
	// can be considered cleanup
	close(closeChan)
	releaseAllResources()
	fmt.Println("The application shut down gracefully")
}

func releaseAllResources() {
	io.WriteString(writer, "Application releasing resources\n")
	writer.Close()
}
