package main

import "fmt"
import "reflect"

func print_string_values(args ...string) {

	for _, e := range(args) {
		fmt.Printf("[%v]", e)
	}
	fmt.Println()
}

func print_float_values(args ...float64) {

	for _, e := range(args) {
		fmt.Printf("[%v]", e)
	}
	fmt.Println()
}

func print_various_values(args ...interface{}) {

	for _, e := range(args) {
		value := reflect.ValueOf(e)
		fmt.Printf("[%v]", value)
	}
	fmt.Println()
}

func main() {
	
	print_string_values("Hello", " ", "World!")
	print_float_values(-982.901, 1, 2, 3, 3.0, 21245, 0.00081)
	print_various_values("Hello", " ", 999, " ", 1.1111, " ", "World!", nil)
}
