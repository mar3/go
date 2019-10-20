package main

import "fmt"

func operation1(args[] string) {

	fmt.Printf("operation1: %v\n", args)
}

func operation2(args[] string) {

	fmt.Printf("operation2: %s\n", args)
}

// 高階関数
// ※引数に関数を含む
func invoke(f func([]string), args []string) {

	f(args)
}

// 高階関数
// ※戻り値が関数
func resolve(id string) func([]string) {

	switch id {
		case "1":
			return operation1
		case "2":
			return operation2
	}

	return nil
}

func call(id string, args []string) {

	operation := resolve(id)
	if operation == nil {
		fmt.Printf("[WARN] invalid request %v.\n", id)
		return
	}

	invoke(operation, args)
}

func main() {

	args := []string{"Hello"}

	call("0", args)
	call("1", args)
	call("2", args)
}
