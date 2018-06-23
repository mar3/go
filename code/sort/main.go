package main

import "sort"
// import "strings"
import "fmt"

func main() {
	
	items := []string{
		"f",
		"J",
		"O",
		"h",
		"l",
		"b",
		"R",
		"F",
		"c",
		"G",
		"z",
		"t",
		"x",
		"a",
		"t",
		"B",
		"e",
		"o",
	}
	sort.Strings(items)
	fmt.Println(items)
}
