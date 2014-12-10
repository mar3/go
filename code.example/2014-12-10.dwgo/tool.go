//
// #DWGO 2014-12-10 at KABUKIZA TOWER
//
// SPECIAL THANKS TO:
//	- yukkurisinai (https://github.com/yukkurisinai)
//	- all the others in #DWGO
//

package main

import (
	"os"
	"log"
	"io"
	"net/http"
	"strings"
	"github.com/codegangsta/cli"
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

func _inner_main(c *cli.Context) {
	// println("Hello my friend!")
	for i, path := range c.Args[1:] {
		tmp := strings.Split(path, "/")
		log.Println(path)
		err := download(path, tmp[len(tmp)-1])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "myget"
	app.Usage = "fight the loneliness!"
	app.Action = _inner_main
	app.Run(os.Args)
}
