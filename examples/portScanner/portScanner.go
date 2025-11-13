package main

import (
	"fmt"
	"net"
)

// To scan ip address open port
// you need target ip address
// and example port

func main() {
	IP := "scanme.nmap.org"
	//Port := "80" // that http port

	// use net/Dial func
	for port := 1; port < 100; port++ {

		address := fmt.Sprintf("%s:%d", IP, port)
		network := "tcp"
		connection, err := net.Dial(network, address)
		if err == nil {
			fmt.Printf("Connection establish... %v [%s]", port, connection.RemoteAddr().String())
		}
	}
}
