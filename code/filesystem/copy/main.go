package main

import (
	"copy/application"
	"fmt"
	"os"
)

func getArg(position int) string {

	if len(os.Args) <= position {
		return ""
	}
	return os.Args[position]
}

func usage() {

	fmt.Println("USAGE:")
	fmt.Println("    copy src dest")
	fmt.Println()
}

func main() {

	left := getArg(1)
	right := getArg(2)
	if left == "" {
		usage()
		return
	}
	if right == "" {
		usage()
		return
	}
	app := application.Application{}
	app.Copy(left, right)
}
