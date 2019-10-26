package main

import "fmt"
import "os"

func main() {

	for _, e := range os.Args[1:] {
		letter := e[0]
		fmt.Printf("%c (%d, 0x%0x)", letter, int(letter), int(letter))
		fmt.Println()
	}
}

