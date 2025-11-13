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
	Port := "80" // that http port

	// use net/Dial func
	address := IP + ":" + Port
	network := "tcp"
	connection, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("Error connection to:", IP)
	}

	fmt.Println("connecion ...", connection.RemoteAddr().String())

}
