package main

import "copy/application"
import "os"

func getArg(position int) string {

	if len(os.Args) <= position {
		return ""
	}
	return os.Args[position]
}

func main() {

	left := getArg(1)
	right := getArg(2)
	if left == "" {
		return
	}
	if right == "" {
		return
	}
	app := application.Application{}
	app.Copy(left, right)
}
