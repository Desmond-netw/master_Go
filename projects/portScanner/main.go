package main

import (
	"fmt"
	"net"
)

func main() {
	// ip
	IP := "scanme.nmap.org"
	Port := "80"

	address := IP + ":" + Port
	network := "tcp"
	connection, err := net.Dial(network, address)

	if err != nil {
		fmt.Println("connection established fail", IP)
	} else {
		fmt.Println("[+] connection established ...", connection.RemoteAddr().String())
	}

}
