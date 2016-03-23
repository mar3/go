package main

import "os"
import "fmt"
import "io/ioutil"

func is_directory(path string) (bool) {

	f, _ := os.Stat(path)
	return f.IsDir()
}

func enumerate_files(path string) {

	if is_directory(path) {
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			enumerate_files(path + "\\" + f.Name())
		}
	} else {
		fmt.Println("DETECTED: " + path)
	}
}

func main() {

	for _, path := range os.Args[1:] {
		enumerate_files(path)
	}
}
