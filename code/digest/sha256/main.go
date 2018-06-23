package main

import "os"
import "fmt"
import "crypto/sha256"

func main() {

	for _, s := range os.Args[1:] {
		data := []byte(s)
		fmt.Printf("%s: %x\n", s, sha256.Sum256(data))
	}
}
