package main

import "fmt"
import "time"

func dump(a *int) {

	go func() {
		*a = 500
		fmt.Printf("[TRACE] dump: [%v:%v]\n", a, *a)
	}()
}

func main() {

	a := 0
	b := 0
	c := 0

	fmt.Printf("[TRACE] {%v}\n", &a)
	fmt.Printf("[TRACE] {%v}\n", &b)
	fmt.Printf("[TRACE] {%v}\n", &c)

	dump(&a)
	dump(&a)
	dump(&b)
	dump(&b)

	time.Sleep(time.Second * 5)
}
