package main

import (
	"copy/application"
	"fmt"
	"os"
)

// コマンドライン引数の取り出し
func getArg(position int) string {

	if len(os.Args) <= position {
		return ""
	}
	return os.Args[position]
}

// 使用方法を表示します。
func usage() {

	fmt.Println("USAGE:")
	fmt.Println("    copy src dest")
	fmt.Println()
}

// エントリーポイントの定義です。
func main() {

	// コマンドライン引数を二つ取ります。
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

	// アプリケーションを実行します。
	app := application.Application{}
	app.Copy(left, right)
}
