package main

import "fmt"

func main() {

	{
		s := "Hello Go World!"
		fmt.Println(s)
	}

	{
		// undefined: s
		fmt.Println(s)
	}
}

