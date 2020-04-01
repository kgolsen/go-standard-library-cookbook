package main

import (
	"fmt"
	"net"
)

func main() {

	// reverse lookup
	addrs, err := net.LookupAddr("8.8.8.8")
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// forward lookup
	ips, err := net.LookupIP("dns.google.com")
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
