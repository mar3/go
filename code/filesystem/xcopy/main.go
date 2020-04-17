package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func copyFile(left string, right string) error {
	fmt.Println(right)
	leftFile, error := os.Open(left)
	if error != nil {
		return error
	}
	defer leftFile.Close()
	rightFile, error := os.Create(right)
	if error != nil {
		return error
	}
	_, error = io.Copy(leftFile, rightFile)
	if false {
		mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
		atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
		os.Chtimes("some-filename", atime, mtime)
	}
	return error
}

func xcopy(sourcePath string, destinationPath string) error {
	source, _ := os.Stat(sourcePath)
	if source == nil {
		fmt.Printf("[WARN] Source entry not found. [%v]\n", sourcePath)
		return nil
	}
	if source.IsDir() {
		file, _ := ioutil.ReadDir(sourcePath)
		for _, e := range file {
			name := e.Name()
			if name == ".svn" {
				continue
			}
			if name == ".git" {
				continue
			}
			if name == "node_modules" {
				continue
			}
			if name == "Debug" {
				continue
			}
			if name == "Release" {
				continue
			}
			error := os.MkdirAll(destinationPath, 0777)
			if error != nil {
				return error
			}
			fullPath := path.Join(sourcePath, name)
			error = xcopy(fullPath, path.Join(destinationPath, name))
			if error != nil {
				fmt.Println("[ERROR] ", error)
				return error
			}
		}
		return nil
	} else {
		stat, _ := os.Stat(destinationPath)
		if stat != nil && stat.IsDir() {
			src := path.Join(sourcePath, source.Name())
			dest := path.Join(destinationPath, source.Name())
			return copyFile(src, dest)
		}
		// sourceFullPath := path.Join(sourcePath, source.Name())
		return copyFile(sourcePath, destinationPath)
	}
}

func usage() {
	fmt.Println("USAGE:")
	fmt.Println("    xcopy left right")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	xcopy(os.Args[1], os.Args[2])
	fmt.Println("Ok.")
}
