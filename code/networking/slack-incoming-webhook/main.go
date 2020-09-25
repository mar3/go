package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

// 現在のタイムスタンプを返します。
func getCurrentTimestamp() string {
	now := time.Now()
	return now.String()[0:23]
}

type Configuration struct {
	WebhookUrl  string
	AccessToken string
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
	self.AccessToken = fmt.Sprintf("%v", dat["access_token"])

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
	conf *Configuration
}

func (self *Application) Configure() *Configuration {

	if self.conf != nil {
		return self.conf
	}

	self.conf = &Configuration{}
	self.conf.Configure()
	return self.conf
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

func (self *Application) UploadFile(channel string, path string) string {

	//curl \
	//  -F file=@cycling.jpeg \
	//  -F "initial_comment=Hello, Leadville" \
	//  -F channels=C0R7MFNJD \
	//  -H "Authorization: Bearer xoxp-123456789" \
	//  https://slack.com/api/files.upload

	conf := self.Configure()
	if conf == nil {
		return ""
	}

	if false {
		content := make(map[string]string)
		contentText, err := json.Marshal(content)
		if err != nil {
			fmt.Printf("[ERROR] コンテンツの加工に失敗しています。理由: [%s]", err.Error())
			return ""
		}
		fmt.Println(contentText)
	}

	// Slack Bot に与えられたアクセストークン
	accessTokenHeader := fmt.Sprintf("Bearer %s", conf.AccessToken)
	fmt.Printf("[TRACE] accessToken: [%s]\n", accessTokenHeader)

	// values := url.Values{}

	// 送信用バッファ
	var buf bytes.Buffer

	// マルチパート書き込み用オブジェクトを作成
	multipartWriter := multipart.NewWriter(&buf)

	// file フィールド
	{
		// ファイルをオープン
		f, err := os.Open(path)
		if err != nil {
			return ""
		}
		defer f.Close()
		streamWriter, error := multipartWriter.CreateFormFile("file", "main.go")
		if error != nil {
			return ""
		}
		if _, error = io.Copy(streamWriter, f); error != nil {
			return ""
		}
	}

	// フィールド:
	{
		multipartWriter.WriteField("initial_comment", "さあうけとるがよい")
	}

	// フィールド: channels
	if true {
		multipartWriter.WriteField("channels", "#notifications")
	}

	multipartWriter.Close()

	url := "https://slack.com/api/files.upload"

	// POST
	client := getTlsClient()
	req, _ := http.NewRequest("POST", url, &buf)
	fmt.Printf("[TRACE] URL: [%s]\n", url)
	fmt.Printf("[TRACE] Content-Type: [%s]\n", multipartWriter.FormDataContentType())
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Authorization", accessTokenHeader)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] リクエストを完了できません。理由: [%s]\n", err)
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの読み込みに失敗しています。理由: [%s]\n", err.Error())
		return ""
	}

	return string(body)
}

// アクセストークンを使用して GitHub のファイルを DELETE します。
func (self *Application) Run() {

	if false {
		content := make(map[string]string)
		content["text"] = "はじめまして ようこそ おじゃまいたします さようなら ごきげんよう"
		response := self.PostSlackWebhook(content)
		if response == "" {
			return
		}
		fmt.Println(response)
	}

	if true {
		// response := self.UploadFile("CNH0XTKJ7", "main.go")
		response := self.UploadFile("notifications", "main.go")
		fmt.Printf("[TRACE] RESPONSE: [%s]\n", response)
	}
}

func main() {

	app := Application{}
	if app.Configure() == nil {
		return
	}

	app.Run()
}
