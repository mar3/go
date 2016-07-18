package main

import "net/http"
import "text/template"
import "time"
import "fmt"

func handler(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	content := make(map[string]string)
	content["now"] = fmt.Sprintf("%v", time.Now())
	t.Execute(writer, content)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
