package main

import (
	"fmt"
	"os"
)

type Error struct {
	description string
}

func (self Error) Error() string {

	return self.description
}

func execute() error {

	xxxxx := os.Getenv("xxxxx")
	if xxxxx == "" {
		return Error{"環境変数が未定義です。"}
	}
	return nil
}

func main() {

	err := execute()
	if err != nil {
		fmt.Printf("[ERROR] 実行時エラーです。要求は中止されました。理由: %v\n", err)
		return
	}

	fmt.Println("Ok.")
}
