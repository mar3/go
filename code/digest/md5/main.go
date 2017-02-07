package main


import "os"
import "fmt"
import "crypto/md5"

func main() {

	for _, s := range os.Args[1:] {
		data := []byte(s)
		fmt.Printf("%x\n", md5.Sum(data))
	}
}


