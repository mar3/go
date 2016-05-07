package main

import "fmt"
import "container/list"
import "reflect"

func main() {

	l := list.New()
	l.PushBack(1)
	l.PushBack("2:文字列")
	l.PushBack(3.33)

	for e := l.Front(); e != nil; e = e.Next() {
		value := e.Value
		fmt.Printf("value: %v, type: %s\n", value, reflect.TypeOf(value))
	}
}


