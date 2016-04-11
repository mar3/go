package main

import "os"
import "fmt"
import "io/ioutil"
import "path"

func _is_directory(path_string string) (bool) {

	if path_string == "" {
		return false
	}
	f, _ := os.Stat(path_string)
	return f.IsDir()
}

func _enumerate_files(path_string string) {

	if _is_directory(path_string) {
		// fmt.Println(path_string)
		files, _ := ioutil.ReadDir(path_string)
		for _, f := range files {
			_enumerate_files(path.Join(path_string, f.Name()))
		}
	} else {
		fmt.Println(path_string)
	}
}

func main() {

	for _, path_string := range os.Args[1:] {
		_enumerate_files(path_string)
	}
}
