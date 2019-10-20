package main

import "fmt"

func キーの有無を調べる() {

	m := make(map[string]string)

	m["0"] = "Zero"
	m["1"] = "One"

	fmt.Println(m["1"])

	// キーの有無を調べる
	_, test := m["2"]

	fmt.Println(test)
}

func main() {

	// キーが string、値は任意の型を許容するマップ
	m := make(map[string]interface{})

	// キーは文字列
	m["氏名"] = "Jimi Hendrix"
	m["住所"] = "350 Monroe Avenue NE, Greenwood Memorial Park, Renton, WA 98056"
	m["年齢"] = 27
	m["電話番号"] = "090 1234 5678"
	m["ABC"] = nil

	fmt.Printf("氏名: [%v]\n", m["氏名"])
	fmt.Printf("住所: [%v]\n", m["住所"])
	fmt.Printf("年齢: [%v]\n", m["年齢"])
	fmt.Printf("電話番号: [%v]\n", m["電話番号"])
	fmt.Printf("ABC: [%v]\n", m["ABC"])
	fmt.Printf("存在しないキー: [%v]\n", m["存在しないキー"])

	delete(m, "ABC")
	delete(m, "ABC")
	delete(m, "ABC")
	delete(m, "電話番号")
	delete(m, "電話番号")
	delete(m, "電話番号")

	fmt.Println(m)
}

