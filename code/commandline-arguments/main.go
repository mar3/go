package main

import "os"
import "fmt"

func main() {

	fmt.Println(len(os.Args))

	for i, e := range os.Args {
		fmt.Print(i)
		fmt.Print("=")
		fmt.Print(e)
		fmt.Println()
	}
}


