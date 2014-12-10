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
	"log"
	"net/http"
	"os"
	"strings"
)
 
func download(url string, saveFileName string) (err error) {
	saveFile, err := os.Create(saveFileName)
	if err != nil {
		return
	}
	defer saveFile.Close()
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
 
	_, err = io.Copy(saveFile, response.Body)
	if err != nil {
		return
	}
	return
}

func main() {
	for i, path := range os.Args {
		if i != 0 {
			tmp := strings.Split(path, "/")
			log.Println(path)
			err := download(path, tmp[len(tmp)-1])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
