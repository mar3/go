package main

import "fmt"

type Employee struct {
	name string;
}

func main() {

	xmap := make(map[string]Employee)
	e := xmap["undefined id"]
	e.name = "おなまえ"

	// 要素が追加されている
	fmt.Println(e)
}
