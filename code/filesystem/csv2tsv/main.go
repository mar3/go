package main

import "os"
import "io"
import "fmt"
import "encoding/csv"

func main() {
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ','
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			break
		}
		for i, e := range(fields) {
			if 0 < i {
				fmt.Print("\t")
			}
			fmt.Print(e)
		}
		fmt.Println()
	}
}

