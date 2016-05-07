package main

import "fmt"

func handler(arg string) {
	fmt.Printf("handler: [%s]\n", arg)
}

// 高階関数
func invoke(f func(string), arg string) {
	f(arg)
}

func main() {
	invoke(handler, "Hello")
}
