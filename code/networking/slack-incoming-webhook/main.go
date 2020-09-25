package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 現在のタイムスタンプを返します。
func getCurrentTimestamp() string {
	now := time.Now()
	return now.String()[0:23]
}

type Configuration struct {
	WebhookUrl string
}

// コンフィギュレーション
func (self *Configuration) Configure() bool {
	buf, err := ioutil.ReadFile(".settings.json")
	if err != nil {
		fmt.Printf("[ERROR] ファイルを開けません。理由: [%s]\n", err.Error())
		return false
	}

	var dat map[string]string
	err = json.Unmarshal([]byte(string(buf)), &dat)
	if err != nil {
		panic(err)
	}
	self.WebhookUrl = fmt.Sprintf("%v", dat["webhook_url"])

	return true
}

func getTlsClient() *http.Client {

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
		// DisableKeepAlives: true,
	}

	return client
}

func get(url string) string {

	client := getTlsClient()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] リクエストに失敗しています。理由: [%s]", err.Error())
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの読み込みに失敗しています。理由: [%s]", err.Error())
		return ""
	}
	return string(body)
}

type Application struct {
	_conf *Configuration
}

func (self *Application) Configure() *Configuration {

	if self._conf != nil {
		return self._conf
	}

	self._conf = &Configuration{}
	self._conf.Configure()
	return self._conf
}

func (self *Application) PostSlackWebhook(content map[string]string) string {

	conf := self.Configure()
	if conf == nil {
		return ""
	}

	contentText, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの加工に失敗しています。理由: [%s]", err.Error())
		return ""
	}

	client := getTlsClient()
	rd := strings.NewReader(string(contentText))
	req, _ := http.NewRequest("POST", conf.WebhookUrl, rd)
	req.Header.Add("Content_Type", "application/json; charset=UTF-8")
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] %s", err)
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの読み込みに失敗しています。理由: [%s]", err.Error())
		return ""
	}

	return string(body)
}

// アクセストークンを使用して GitHub のファイルを DELETE します。
func (self *Application) Run() {

	content := make(map[string]string)
	content["text"] = "はじめまして ようこそ おじゃまいたします さようなら ごきげんよう"

	response := self.PostSlackWebhook(content)
	if response == "" {
		return
	}

	fmt.Println(response)
}

func main() {

	app := Application{}
	if app.Configure() == nil {
		return
	}

	app.Run()
}
