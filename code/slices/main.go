package main

import "fmt"

func main() {

	l := make([]string, 0)
	fmt.Printf("%v (%d)\n", l, len(l))

	l = append(l, "あ")
	l = append(l, "い")
	l = append(l, "う")
	fmt.Printf("%v (%d)\n", l, len(l))

	fmt.Printf("%v\n", l[1])
}

