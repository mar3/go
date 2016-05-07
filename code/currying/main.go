package main

import "os"
import "fmt"

func getLogger(out *os.File) func(message string) {

	return func(message string) {
		fmt.Fprintf(out, "%s\n", message)
	}
}

func main() {

	log := getLogger(os.Stdout)
	log("hello 1")
	log("hello 2")
	log("hello 3")
}
