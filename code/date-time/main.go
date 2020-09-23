package main

import (
	"fmt"
	"time"
)

// 現在のタイムスタンプを返します。
func getCurrentTimestamp() string {

	now := time.Now()
	return now.String()[0:23]
}

// エントリーポイント
func main() {

	fmt.Printf("current timestamp: [%s]\n", getCurrentTimestamp())
	fmt.Printf("now: [%v]\n", time.Now())
}
