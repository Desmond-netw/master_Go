# Go IP Scanner

A simple command-line **IP/Port scanner** built with **Go** as part of my journey to master **GoLang for Cybersecurity**.

This project demonstrates how to establish TCP connections using Go’s standard `net` package — the same core concept behind many network reconnaissance tools.

---

## Features

- Takes a target IP or hostname.
- Attempts a TCP connection to a given port.
- Prints connection success or failure.
- Lightweight, fast, and beginner-friendly.

---

## Learning Highlights

While building this project, I learned how to:

- Use Go’s `net.Dial()` function to establish network connections.
- Handle network errors gracefully.
- Combine IP addresses and ports into proper network addresses.
- Understand basic concepts of port scanning and network probing.

---

## Demo

Click below to watch a short demo video of the scanner in action 🎥

[![Watch the video](https://img.youtube.com/vi/RrnpLb4qgWc/0.jpg)](https://www.youtube.com/watch?v=RrnpLb4qgWc)

---

## Code Example

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	IP := "scanme.nmap.org" // example safe target for testing
	Port := "80"

	network := "tcp"
	address := IP + ":" + Port
	connection, err := net.Dial(network, address)

	if err != nil {
		fmt.Println("Connection failed:", IP)
		return
	}
	fmt.Println("Connection established:", connection.RemoteAddr().String())
}
