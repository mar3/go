package main

import "fmt"

func main() {

	m := make(map[string]interface{})
	m["氏名"] = "Jimi Hendrix"
	m["住所"] = "350 Monroe Avenue NE, Greenwood Memorial Park, Renton, WA 98056"
	m["年齢"] = 27
	m["電話番号"] = ""

	fmt.Println(m)
}

