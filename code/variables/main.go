package main

import "fmt"

func main() {

	{
		const d = 9.999
		// d = 200 #cannot assign to d
		fmt.Println(d)
	}

	{
		var s string = "Hello Go World!!"
		fmt.Println(s)
	}

	{
		s := "Hello Go World!!"
		fmt.Println(s)
	}
}
