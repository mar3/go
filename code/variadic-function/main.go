package main

import "fmt"

//
// string のみ
//
func println_01(args ...string) {

	for _, e := range(args) {
		fmt.Print(e)
	}
	fmt.Println()
}

//
// int のみ
//
func println_02(args ...int) {

	for _, e := range(args) {
		fmt.Print(e)
	}
	fmt.Println()
}

//
// 様々な型
//
func println_03(args ...interface{}) {

	for _, e := range(args) {
		fmt.Print(e)
	}
	fmt.Println()
}

func main() {
	
	println_01("Hello", " ", "World!")
	println_02(1, 2, 3, 4, 5)
	println_03("Hello", " ", 999, " ", 1.1111, " ", "World!")
}
