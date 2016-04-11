package main

import "fmt"
// import "strings"

type user_information struct {
	id string
	name string
	age int
}

func main() {

	user := user_information{
		id: "9182528-17-3141",
		name: "Wynton Kelly"}

	fmt.Println(user)
}
