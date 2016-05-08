package main

import "fmt"

type User struct {
	id int
	name string
	email string
}

func (this *User) toJSON() string {
	return fmt.Sprintf("{id:%v, name:\"%v\", email:\"%v\"}",
		this.id, this.name, this.email)
}

func main() {
	user := User{id: 1, name: "Max Middleton", email: "david.max.middleton@gmail.com"}
	fmt.Println(user.toJSON())
}
