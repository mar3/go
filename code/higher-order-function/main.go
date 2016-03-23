//
// 高階関数
//

package main

import "fmt"

func handler(arg string) {
	fmt.Print("handler: ")
	fmt.Println(arg)
}

// run() は“高階関数”(=higher-order function)である
func run(f func(string), arg string) {
	f(arg)
}

func main() {
	run(handler, "Hello")
}
