package main

import "fmt"
import "os"
import "github.com/google/uuid"
import "io/ioutil"
import "path/filepath"

func generateNewName() (string) {
	id := uuid.New()
	return fmt.Sprintf("%v", id)
}

func rename(path string) {
	parent := filepath.Dir(path)
	ext := filepath.Ext(path)
	newpath := filepath.Join(parent, generateNewName() + ext)
	fmt.Printf("rename ... [%v] >> [%v]", path, newpath)
	fmt.Println()
	os.Rename(path, newpath)
}

func run(s string, operation func(string)) {
	info, err := os.Stat(s)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		fmt.Println()
		return
	}
	if info.IsDir() {
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
