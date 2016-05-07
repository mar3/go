package main

import "os"
import "bufio"
import "fmt"

func main() {

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}

