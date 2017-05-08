package main


import "fmt"
import "bufio"
import "strings"


func main() {

	s := "ああああああああああ\r\nあああああああああああ\nいいいいいいいい\nうううううううううう"
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("[" + line + "]")
	}
}

