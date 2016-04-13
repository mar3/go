package main

import "os"
// import "fmt"
import "io/ioutil"
import "path"
import "log"
import "strings"
import "os/exec"
import "bytes"
// import "io"

// func _save(p string, content string) {

// }

func _decompile(path_string string) {

	if !strings.HasSuffix(path_string, ".class") {
		return
	}

	command := exec.Command(
		"java", "-jar", "C:\\CFR\\cfr_0_115.jar", path_string)
	var buffer bytes.Buffer
	command.Stdout = &buffer
	error := command.Run()
	if error != nil {
		log.Fatal(error)
		return
	}

	stream, error := os.Create(path_string + ".java")
	if error != nil {
		log.Fatal(error)
		return
	}
	defer stream.Close()
	length, error := stream.Write(buffer.Bytes())
	if error != nil {
		log.Fatal(error)
		return
	}
	if length == 0 {

	}
}

func _is_directory(path_string string) (bool) {

	if path_string == "" {
		return false
	}
	f, _ := os.Stat(path_string)
	return f.IsDir()
}

func _enumerate_files(path_string string) {

	if _is_directory(path_string) {
		files, _ := ioutil.ReadDir(path_string)
		for _, f := range files {
			_enumerate_files(path.Join(path_string, f.Name()))
		}
	} else {
		_decompile(path_string)
	}
}

func main() {

	for _, path_string := range os.Args[1:] {
		_enumerate_files(path_string)
	}
}
