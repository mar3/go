package main

import "fmt"
import "flag"

func main() {

	boolean_value := flag.Bool("help", false, "boolean value")
	flag.Parse()
	fmt.Printf("boolean value: [%v]\n", *boolean_value)
}

