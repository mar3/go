package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

// 現在のタイムスタンプを返します。
func getCurrentTimestamp() string {
	now := time.Now()
	return now.String()[0:23]
}

// コンフィギュレーション
func configure() string {
	buf, err := ioutil.ReadFile(".access-token.txt")
	if err != nil {
		return ""
	}
	return string(buf)
}

// 送信するファイルのコンテンツを返します。
func createFileContent() []byte {
	content := fmt.Sprintf("UPDATE: %s", getCurrentTimestamp())
	return []byte(content)
}

// アクセストークンを使用して GitHub のファイルを DELETE します。
func deleteFile(accessToken string) {

	c := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(c, tokenSource)
	client := github.NewClient(httpClient)

	fileContent := createFileContent()

	committer := github.CommitAuthor{
		Name:  github.String("Masaru Irisawa"),
		Email: github.String("mass10.github@gmail.com"),
	}

	// ========== GET REPOSITORY CONTENT ==========
	{
		options := &github.RepositoryContentGetOptions{}
		content, _, _, err := client.Repositories.GetContents(c, "mass10", "DEPLOY-REPOSITORY", "example-01", options)
		if err != nil {
			fmt.Printf("[ERROR] %s", err)
			return
		}
		var sha = *content.SHA
		fmt.Printf("[TRACE] content.SHA: [%s]\n", sha)
	}

	return

	options := &github.RepositoryContentFileOptions{
		Message:   github.String("COMMIT MESSAGE HERE"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &committer,
	}
	// 	curl \
	//   -X DELETE \
	//   -H "Accept: application/vnd.github.v3+json" \
	//   https://api.github.com/repos/octocat/hello-world/contents/PATH \
	//   -d '{"message":"message","sha":"sha"}'
	repo, status, err := client.Repositories.DeleteFile(c, "mass10", "DEPLOY-REPOSITORY", "example-02", options)
	if err != nil {
		fmt.Printf("[ERROR] %s", err)
		return
	}
	if status != nil {
		fmt.Printf("[TRACE] STATUS: %v\n", status.Body)
	}
	if repo != nil {
		fmt.Printf("[TRACE] REPO: %v\n", repo)
	}
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

// アクセストークンを使用して GitHub のファイルを DELETE します。
func getRepositoryContent(accessToken string) {

	// ========== GET REPOSITORY CONTENT ==========
	{
		client := getTlsClient()
		url := "https://api.github.com/repos/mass10/DEPLOY-REPOSITORY/contents/README.md"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Accept", "application/vnd.github.v3+json")
		req.Header.Add("Authorization", fmt.Sprintf("token %s", accessToken))
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
		fmt.Println("")
		fmt.Println("---")
		fmt.Println(string(body))
		fmt.Println("---")
		fmt.Println("")

		var dat map[string]interface{}
		err = json.Unmarshal([]byte(string(body)), &dat)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[TRACE] name: [%s]\n", dat["name"])
		fmt.Printf("[TRACE] path: [%s]\n", dat["path"])
		fmt.Printf("[TRACE] sha: [%s]\n", dat["sha"])
		fmt.Printf("[TRACE] size: [%f]\n", dat["size"])
		fmt.Printf("[TRACE] type: [%s]\n", dat["type"])
	}
}

// アクセストークンを使用して GitHub にファイルを PUSH します。
func createCommitPush(accessToken string) {

	c := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(c, tokenSource)
	client := github.NewClient(httpClient)

	fileContent := createFileContent()

	committer := github.CommitAuthor{
		Name:  github.String("Masaru Irisawa"),
		Email: github.String("mass10.github@gmail.com"),
	}

	options := &github.RepositoryContentFileOptions{
		Message:   github.String("COMMIT MESSAGE HERE"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &committer,
	}

	repo, status, err := client.Repositories.CreateFile(c, "mass10", "DEPLOY-REPOSITORY", "example-02", options)
	if err != nil {
		fmt.Printf("[ERROR] %s", err)
		return
	}
	if status != nil {
		fmt.Printf("[TRACE] STATUS: %v\n", status.Body)
	}
	if repo != nil {
		fmt.Printf("[TRACE] REPO: %v\n", repo)
	}
}

// アクセストークンを使用して GitHub にファイルを PUSH します。
func createCommitPushNew(accessToken string) {

	// ========== セッションを初期化 ==========
	c := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(c, tokenSource)
	client := github.NewClient(httpClient)

	sha := ""
	userName := "mass10"
	repositoryName := "DEPLOY-REPOSITORY"
	contentName := "example-01"

	// ========== GET REPOSITORY CONTENT ==========
	{
		options := &github.RepositoryContentGetOptions{}
		content, _, _, err := client.Repositories.GetContents(c, userName, repositoryName, contentName, options)
		if err != nil {
			fmt.Printf("[ERROR] コンテンツを参照できませんでした。理由: [%s]", err)
			return
		}
		if content == nil {
			fmt.Printf("[TRACE] コンテンツを参照できませんでした。理由: content が nil です。")
		}
		sha = *content.SHA
		fmt.Printf("[TRACE] content.SHA: [%s]\n", sha)
	}

	// ========== PUSH CONTENT ==========
	{
		// Push するファイルの内容
		fileContent := createFileContent()
		committer := github.CommitAuthor{
			Name:  github.String("Masaru Irisawa"),
			Email: github.String("mass10.github@gmail.com"),
		}
		options := &github.RepositoryContentFileOptions{
			Message:   github.String("COMMIT MESSAGE HERE"),
			Content:   fileContent,
			Branch:    github.String("master"),
			Committer: &committer,
		}
		if sha != "" {
			// ファイルを更新する場合
			fmt.Println("[TRACE] ファイルが存在するため、SHA を設定しています。")
			options.SHA = github.String(sha)
		} else {
			fmt.Println("[TRACE] ファイルは存在しません。SHA は空です。")
		}
		repo, status, err := client.Repositories.CreateFile(c, userName, repositoryName, contentName, options)
		if err != nil {
			fmt.Printf("[ERROR] ファイルを更新できません。理由: [%s]", err)
			return
		}
		if status != nil {
			fmt.Printf("[TRACE] STATUS: %v\n", status.Body)
		}
		if repo != nil {
			fmt.Printf("[TRACE] REPO: %v\n", repo)
		}
	}

	fmt.Println("[TRACE] SUCCESS!!")
}

func main() {

	accessToken := configure()

	// deleteFile(accessToken)

	// Pure Go
	getRepositoryContent(accessToken)

	// createCommitPush(accessToken)
	// createCommitPushNew(accessToken)
}
