package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/tls"
)

func getNormalClient() (*http.Client) {

	client := &http.Client{}
	return client
}

func getTlsClient() (*http.Client) {

	conf := &tls.Config{
		// InsecureSkipVerify: false,
		// ServerName: "google.co.jp",
	}

	tr := &http.Transport{
		TLSClientConfig: conf,
		// Proxy: http.ProxyFromEnvironment,
	}

	client := &http.Client{
		Transport: tr,
		DisableKeepAlives: true,
	}

	return client
}

func main() {

	if false {
		os.Setenv("HTTP_PROXY", "http://proxy.company.co.jp")
	}

	// client := getNormalClient()
	client := getTlsClient()

	url := "https://www.google.co.jp/search?q=datadog"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	response, err := client.Do(req)
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

