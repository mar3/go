package main

import "fmt"
import "os"
import "io/ioutil"
import "net/http"
import "encoding/xml"
import "encoding/json"
import "strings"

type Item struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

type Channel struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	Language string `xml:"language"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

func _parse_xml(content []byte) Rss {

	rss := Rss{}
	xml.Unmarshal(content, &rss)
	// fmt.Println("title: ["rss)
	return rss
}

func get_content(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return nil
	}
	return body
}

func read_file(filepath string) []byte {

	file, _ := os.Open(filepath)
	defer file.Close()
	body, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return nil
	}
	return body
}

func save(rss Rss) {

	json, err := json.Marshal(rss)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(json))
}

func load_xml(url string) {

	content := []byte{}
	if strings.HasPrefix(url, "http") {
		content = get_content(url)
		rss := _parse_xml(content)
		save(rss)
	} else {
		content = read_file(url)
		rss := _parse_xml(content)
		save(rss)
	}
}

func main() {

	url := ""
	url = "http://feedblog.ameba.jp/rss/ameblo/passpo-staff/rss20.xml"
	// url = "sample.xml"
	load_xml(url)
}
