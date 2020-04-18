package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Stopwatch struct {
	time time.Time
}

func Stopwatch_New() *Stopwatch {
	watch := Stopwatch{time: time.Now()}
	return &watch
}

func (self *Stopwatch) Reset() {
	self.time = time.Now()
}

func (self *Stopwatch) ToString() string {
	currentDuration := time.Since(self.time)
	millisec := currentDuration.Milliseconds()
	hours := millisec / 1000 / 60 / 60
	minutes := millisec / 1000 / 60
	secs := millisec / 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours,
		minutes,
		secs,
		millisec%1000)
}

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
	// TODO: タイムスタンプの貼り付け
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
		return copyFile(sourcePath, destinationPath)
	}
}

func usage() {
	fmt.Println("USAGE:")
	fmt.Println("    backup target")
}

func getTargetPath(p string) string {
	now := time.Now()
	return fmt.Sprintf("%s-%04d%02d%02d-%02d%02d%02d",
		p, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func backup(left string) {
	leftStat, err := os.Stat(left)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	if leftStat.IsDir() {
		leftAbsolutePath, err := filepath.Abs(left)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err)
			return
		}
		rightAbsolutePath := getTargetPath(leftAbsolutePath)
		watch := Stopwatch_New()
		xcopy(left, rightAbsolutePath)
		fmt.Println("Ok. (elaped: " + watch.ToString() + ")")
	} else {
		fmt.Println("ディレクトリを指定してください。")
	}
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	backup(os.Args[1])
}
