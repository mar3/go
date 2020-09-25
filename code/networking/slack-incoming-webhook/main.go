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

// webhook を利用してメッセージを投稿します。
// コンソールで設定されたチャネルに投稿されます。
func (self *Application) PostMessageByWebhook(text string) {

	conf := self.Configure()
	if conf == nil {
		return
	}

	content := make(map[string]string)
	content["text"] = text
	contentText, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの加工に失敗しています。理由: [%s]", err.Error())
		return
	}

	client := getTlsClient()
	rd := strings.NewReader(string(contentText))
	req, _ := http.NewRequest("POST", conf.WebhookUrl, rd)
	req.Header.Add("Content_Type", "application/json; charset=UTF-8")
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] %s", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの読み込みに失敗しています。理由: [%s]", err.Error())
		return
	}

	fmt.Println(string(body))
}

// ファイルの名前部分を返します。
func getFileName(path string) string {
	file, error := os.Stat(path)
	if error != nil {
		return ""
	}
	return file.Name()
}

func (self *Application) TestUploadFile(channel string, path string) {

	conf := self.Configure()
	if conf == nil {
		return
	}

	name := getFileName(path)
	if name == "" {
		fmt.Println("[ERROR] ファイル名が不明です。")
		return
	}

	if false {
		content := make(map[string]string)
		contentText, err := json.Marshal(content)
		if err != nil {
			fmt.Printf("[ERROR] コンテンツの加工に失敗しています。理由: [%s]", err.Error())
			return
		}
		fmt.Println(contentText)
	}

	// Slack Bot に与えられたアクセストークン
	accessTokenHeader := fmt.Sprintf("Bearer %s", conf.AccessToken)
	fmt.Printf("[TRACE] accessToken: [%s]\n", accessTokenHeader)
	// 送信用バッファ
	var buf bytes.Buffer
	// マルチパート書き込み用オブジェクトを作成
	multipartWriter := multipart.NewWriter(&buf)

	// file フィールド
	{
		// ファイルをオープン
		f, err := os.Open(path)
		if err != nil {
			return
		}
		defer f.Close()
		streamWriter, error := multipartWriter.CreateFormFile("file", name)
		if error != nil {
			return
		}
		if _, error = io.Copy(streamWriter, f); error != nil {
			return
		}
	}

	multipartWriter.WriteField("initial_comment", "さあうけとるがよい")
	multipartWriter.WriteField("channels", "#notifications")
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
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] コンテンツの読み込みに失敗しています。理由: [%s]\n", err.Error())
		return
	}

	fmt.Println(string(body))
}

// アクセストークンを使用して GitHub のファイルを DELETE します。
func (self *Application) Run() {

	if false {
		self.PostMessageByWebhook("はじめまして ようこそ おじゃまいたします さようなら ごきげんよう")
	}
	if true {
		self.TestUploadFile("notifications", "cb8ba27acf90f647d362a851a311d801_1428748299.jpg")
	}
}

func main() {

	app := Application{}
	if app.Configure() == nil {
		return
	}

	app.Run()
}
