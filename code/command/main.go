package main

import "fmt"
import "os/exec"
import "log"

func main() {

	command := exec.Command("ls", "-lF", "/tmp")
	out, error := command.Output()
	if error != nil {
		log.Fatal(error)
		return
	}
	fmt.Println(string(out))
}

