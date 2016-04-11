package main

import "fmt"

func main() {

	list := []string {"a", "b", "c"}

	for _, s := range list {
		fmt.Println(s)
	}
}
