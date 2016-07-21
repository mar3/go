package main

import "net/http"
import "text/template"
import "time"
import "fmt"

// [/hello] に対するハンドラーを定義します。
func helloHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("hello"))
}

// [/api] に対するハンドラーを定義します。
func apiHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("{\"key-1\": \"value-1\", \"key-2\": \"value-2\"}"))
}

// [/] に対するハンドラーを定義します。
func defaultHandler(writer http.ResponseWriter, request *http.Request) {

	t, _ := template.ParseFiles("index.html")
	content := make(map[string]string)
	content["now"] = fmt.Sprintf("%v", time.Now())
	content["you"] = "名無し"
	t.Execute(writer, content)
}

func main() {

	// URL のマッピングを定義
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/", defaultHandler)
	// サーバーを起動
	http.ListenAndServe("0.0.0.0:8080", nil)
}
