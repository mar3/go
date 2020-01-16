package main

import "os"
import "bufio"
import "fmt"

func prompt(text string) string {

	fmt.Print(text)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		return line
	}
	return ""
}

func main() {

	fmt.Println("### START ###")
	for {
		answer := prompt("何か入力してください。\n> ")
		fmt.Printf("ANSWER: [%v]\n", answer)
	}
	fmt.Println("Ok.")
}
