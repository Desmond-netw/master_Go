package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	// initilaize a command
	command := "lsae"

	// get the command exc object
	cmd_obj := exec.Command(command)
	cmd_obj.Stdout = os.Stdout

	// edge cases
	cmd_obj.Stderr = os.Stderr
	err := cmd_obj.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
