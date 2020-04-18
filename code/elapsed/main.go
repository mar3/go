package main

import (
	"elapsed/stopwatch"
	"fmt"
	"time"
)

func main() {

	fmt.Println("### start ###")
	watch := stopwatch.New()
	time.Sleep(time.Millisecond * 3456)
	fmt.Printf("ELAPSED: [%s]\n", watch.ToString())
	fmt.Println("--- end ---")
	fmt.Println("Ok.")
}
