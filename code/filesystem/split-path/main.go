package main

import "os"
import "fmt"
import "path/filepath"

func main() {

	for _, e := range os.Args[1:] {

		fmt.Printf("Dir=[%v], Base=[%v], Ext=[%v]\n",
			filepath.Dir(e),
			filepath.Base(e),
			filepath.Ext(e))
		// fmt.Println(e)
	}
}

