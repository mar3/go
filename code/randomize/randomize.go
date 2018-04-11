package main

import "fmt"
import "os"
import "github.com/google/uuid"
import "io/ioutil"
import "path/filepath"

func isDir(name string) bool {

	info, err := os.Stat(name)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func generateNewName() (string) {

	id := uuid.New()
	return fmt.Sprintf("%v", id)
}

func rename(path string) {

	parent := filepath.Dir(path)
	// name := filepath.Base(path)
	ext := filepath.Ext(path)

	newpath := filepath.Join(parent, generateNewName() + ext)
	fmt.Printf("rename ... [%v] >> [%v]\n", path, newpath)
	os.Rename(path, newpath)
}

func run(s string, operation func(string)) {

	if s == "" {
		return
	}
	if isDir(s) {
		entries, _ := ioutil.ReadDir(s)
		for _, f := range entries {
			path := filepath.Join(s, f.Name())
			run(path, operation)
		}
	} else {
		operation(s)
	}
}

func main() {

	for _, e := range(os.Args[1:]) {
		run(e, rename)
	}
}
