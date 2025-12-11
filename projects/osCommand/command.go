// package main
package main

// imports
import (
	"log"
	"os"
	"os/exec"
)

// a function to exe the command
func executeCommand(command string, argArr []string) (err error) {

	arg := argArr
	// creating a command obj
	cmd_obj := exec.Command(command, arg...)
	cmd_obj.Stdout = os.Stdout // standard output
	// handle stderr
	cmd_obj.Stderr = os.Stderr
	err = cmd_obj.Run() //edge cases
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return nil
}

func main() {
	// initilize command
	command := "ls"
	arg := []string{"-ls"}
	//call exec func
	executeCommand(command, arg)
}
