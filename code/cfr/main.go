package main

import "os"
import "io/ioutil"
import "path"
import "log"
import "strings"
import "os/exec"
import "bytes"

func decompile(path_string string) (int) {

	if !strings.HasSuffix(path_string, ".class") {
		return 0
	}

	command := exec.Command(
		"java", "-jar", "C:\\CFR\\cfr_0_115.jar", path_string)
	var buffer bytes.Buffer
	command.Stdout = &buffer
	error := command.Run()
	if error != nil {
		log.Fatal(error)
		return 0
	}

	stream, error := os.Create(path_string + ".java")
	if error != nil {
		log.Fatal(error)
		return 0
	}
	defer stream.Close()
	length, error := stream.Write(buffer.Bytes())
	if error != nil {
		log.Fatal(error)
		return 0
	}
	return length
}

func isDirectory(path_string string) (bool) {

	if path_string == "" {
		return false
	}
	f, _ := os.Stat(path_string)
	return f.IsDir()
}

func enumerateFiles(path_string string) {

	if isDirectory(path_string) {
		files, _ := ioutil.ReadDir(path_string)
		for _, f := range files {
			enumerateFiles(path.Join(path_string, f.Name()))
		}
	} else {
		decompile(path_string)
	}
}

func main() {

	for _, path_string := range os.Args[1:] {
		enumerateFiles(path_string)
	}
}
