package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// なにもしないスキャン
func Scan0() {

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}

// 文字列を Shift_JIS → UTF-8 変換
func Encode(b []byte) string {

	buffer := bytes.NewBuffer(b)
	line, err := ioutil.ReadAll(transform.NewReader(buffer, japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		log.Fatal(err)
	}
	return string(line)
}

// 1行ごとに Shift_JIS → UTF-8 変換
func Scan1() {

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := Encode(s.Bytes())
		fmt.Println(line)
	}
}

// 標準入力を Shift_JIS → UTF-8 変換
func Scan3() {

	in := transform.NewReader(os.Stdin, japanese.ShiftJIS.NewDecoder())
	s := bufio.NewScanner(in)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

// エントリーポイント
func main() {

	// Scan0()
	// Scan1()
	Scan3()
}
