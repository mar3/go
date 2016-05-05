package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
 
func main() {
	response, err := http.Get("http://www.yahoo.co.jp")
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return
	}
	fmt.Println(string(body))
}
