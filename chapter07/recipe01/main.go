package main

import (
	"fmt"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println(iface.Name)
		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}

}
