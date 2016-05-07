package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	p := point{255, 128}
	fmt.Printf("point: %v, (type: %T)\n", p, p)
	fmt.Printf("%t or %t\n", true, false)
	fmt.Printf("%d, %c, %f\n", 999, 65, 1.111111111111)
	fmt.Printf("%b, %x\n", 6, 255)
	fmt.Printf("%q\n", "12123-34324-5-446461")
	fmt.Printf("%s\n", "12123-34324-5-446461")
	fmt.Printf("%x\n", "Hello")
	fmt.Printf("%x, %x\n", "あ", "あ")
}
