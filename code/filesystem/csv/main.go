package main

import "os"
import "io"
import "fmt"
import "encoding/csv"

func main() {

	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("%#v\n", fields)
	}
}
