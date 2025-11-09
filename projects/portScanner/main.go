// a network port scanner app
package main

import (
	"fmt"
	"net"
)

func main() {
	// define static ip to scan
	IP := "scanme.nmap.org" // example: ip address
	Port := "80"

	// to establish connection, let use net package and Dial func to
	// receive ack from the target IP address

	// dial func uses network and the target address:port
	network := "tcp"
	address := IP + ":" + Port // ip:port number
	connection, err := net.Dial(network, address)

	// let handle some edgecases
	if err != nil {
		fmt.Println("connection establishing failed", IP)
		return
	} else {
		fmt.Println("Connection established ...", connection.RemoteAddr().String()) // this will give us the ip we need

	}
}
