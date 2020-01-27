package util

import "time"

// タイムスタンプを返します。
func GetTimestamp() string {

	now := time.Now()
	return now.String()[0:23]
}
