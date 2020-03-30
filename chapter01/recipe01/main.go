package main

import (
	"log"
	"runtime"
)

const (
	info = `
The application %s is starting.
Built with Go version: %s`
)

func main() {
	log.Printf(info, "Example", runtime.Version())
}
