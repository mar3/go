package main

import "fmt"
import "reflect"

type Employee struct {
	name string;
}

func enumEmployees() map[string]Employee {

	employees := make(map[string]Employee)
	return employees
}

func main() {

	employees := enumEmployees()

	// 無いキーを指定しても main.Employee{} が帰ってくる
	e := employees["いない人"]
	fmt.Printf("[TRACE] element is %v {%v}", e, reflect.TypeOf(e))
	fmt.Println()

	// 元のオブジェクトは変化しない
	fmt.Println(employees)
}
