package main

import "fmt"
import "strings"

func main() {

	elements := strings.Split("AAA,BBB,CCC", ",")
	for index, e := range elements {
		fmt.Printf("%d: [%s]\n", index, e)
	}
}
