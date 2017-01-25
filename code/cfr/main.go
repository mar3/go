package main

import "os"
import "fmt"
import "io/ioutil"
import "path"
import "log"
import "strings"
import "os/exec"
import "bytes"

func decompile(path_string string) bool {

	if !strings.HasSuffix(path_string, ".class") {
		return false
	}

	command := exec.Command(
		"java", "-jar", "C:\\CFR\\cfr_0_115.jar", path_string)
	var buffer bytes.Buffer
	command.Stdout = &buffer
	error := command.Run()
	if error != nil {
		log.Fatal(error)
		return false
	}

	stream, error := os.Create(path_string + ".java")
	if error != nil {
		log.Fatal(error)
		return false
	}
	defer stream.Close()
	_, error = stream.Write(buffer.Bytes())
	if error != nil {
		log.Fatal(error)
		return false
	}
	return true
}

func isDirectory(path_string string) (bool) {

	if path_string == "" {
		return false
	}
	f, _ := os.Stat(path_string)
	return f.IsDir()
}

func enumerateFiles(path_string string, ff func(string) bool) bool {

	if isDirectory(path_string) {
		files, _ := ioutil.ReadDir(path_string)
		for _, f := range files {
			pathname := path.Join(path_string, f.Name())
			if enumerateFiles(pathname, ff) {
				return true
			}
		}
		return false
	} else {
		return ff(path_string)
	}
}

func detect_cfr(path string) string {

	path_to_cfr := ""
	ff := func(pathname string) bool {
		if strings.HasPrefix(pathname, "cfr") == false {
			return false
		}
		if strings.HasSuffix(pathname, ".jar") == false {
			return false
		}
		fmt.Println("cfr [" + pathname + "] が検出されました。")
		path_to_cfr = pathname
		return true
	}
	enumerateFiles(path, ff)
	return path_to_cfr
}

func main() {

	path_to_cfr := detect_cfr(".")
	if path_to_cfr == "" {
		fmt.Println("no cfr.")
		return
	}

	// fmt.Println("found [" + path_to_cfr + "]")

	for _, path_string := range os.Args[1:] {
		enumerateFiles(path_string, decompile)
	}
}
