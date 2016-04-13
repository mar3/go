package main

import "fmt"
import "os/exec"
import "log"
import "bytes"

func main() {

	command := exec.Command("ls", "-lF", "/tmp")
	var buffer bytes.Buffer
	command.Stdout = &buffer
	error := command.Run()
	if error != nil {
		log.Fatal(error)
		return
	}
	fmt.Println(buffer.String())
}

