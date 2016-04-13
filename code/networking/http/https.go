package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/tls"
)
 
func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	client := &http.Client{
		Transport: tr,
	}

	url := "https://www.google.co.jp"
	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response_body == nil {
		return
	}

	fmt.Println(string(response_body))
}

