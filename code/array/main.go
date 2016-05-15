package main


import "fmt"
import "reflect"

func main() {

	int_array := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(int_array))
	fmt.Println(cap(int_array))
	fmt.Println(len(int_array))
}

