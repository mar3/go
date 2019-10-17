package main

import "fmt"
import "github.com/google/uuid"

func generateUUID() (string) {
	id := uuid.New()
	return fmt.Sprintf("%v", id)
}

func main() {
	uuid := generateUUID()
	fmt.Println(uuid)
}
