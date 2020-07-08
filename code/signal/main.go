// signal を受け取って安全に処理を中止するバッチ処理の例 [WIP]

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func sub(c <-chan os.Signal, handler chan string) {

	for {
		fmt.Println("[TRACE] <sub> (running...)")
		select {
		case response := <-c:
			fmt.Printf("[TRACE] <goroutine> got message [%s]. exit.\n", response)
			handler <- "exit"
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	fmt.Println("[TRACE] <main> ### START ###")
	c := make(chan os.Signal, 1)
	fmt.Println("[TRACE] <main> (notify)")

	signal.Notify(c, os.Interrupt)
	fmt.Println("[TRACE] <main> (waiting...)")

	handler := make(chan string)
	go sub(c, handler)

	response := <-handler
	fmt.Printf("[TRACE] <main> RECV: [%v]\n", response)
	fmt.Println("[TRACE] <main> --- END ---")
}
