package main

import "os"
import "fmt"

func get_logger(out *os.File) func(message string) {

	return func(message string) {
		fmt.Fprintf(out, "%s\n", message)
	}
}

func main() {

	logger := get_logger(os.Stdout)

	logger("hello")
	logger("hello")
	logger("hello")
}
