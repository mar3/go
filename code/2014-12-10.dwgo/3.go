//
// #DWGO 2014-12-10 at KABUKIZA TOWER
//
// SPECIAL THANKS TO:
//	- yukkurisinai (https://github.com/yukkurisinai)
//	- all the others in #DWGO
//

package main
 
import (
  "io"
  // "io/ioutil"
  "net/http"
  "os"
  "sync"
  "strings"
  // "log"
)
 
func download(url string, saveFileName string) {
	saveFile, _ := os.Create(saveFileName)
	defer saveFile.Close()
	response, _ := http.Get(url)
	defer response.Body.Close()
	io.Copy(saveFile, response.Body)
}
 
func main() {
	var waitGroup sync.WaitGroup
	for i, path := range os.Args[1:] {
		waitGroup.Add(1)
		path_array := strings.Split(path, "/")
		file_name := path_array[len(path_array)-1]
		go func(function func()) {
			defer waitGroup.Done()
			function()
		}(download(path, file_name))
	}
	waitGroup.Wait()
}