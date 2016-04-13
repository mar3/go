package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
 
func main() {
	response, err := http.Get("http://www.google.co.jp")
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

