package main

import "os"
import "fmt"
import "io/ioutil"
import "path"

func enumerateFiles(path_string string) {

	if path_string == "" {
		return
	}

	f, _ := os.Stat(path_string)
	if f.IsDir() {
		files, _ := ioutil.ReadDir(path_string)
		for _, f := range files {
			enumerateFiles(path.Join(path_string, f.Name()))
		}
	} else {
		fmt.Println(path_string)
	}
}

func main() {

	for _, path_string := range os.Args[1:] {
		enumerateFiles(path_string)
	}
}
