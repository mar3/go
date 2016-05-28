package main

import "fmt"

type S struct {
	id int
}

func main() {

	s := S{9999}
	fmt.Printf("address: %p, value: %v, type: %T\n", &s, s, s);
}

