package main

import "os"
import "fmt"
import "bytes"
import "text/template"

func test1() {
	t, _ := template.ParseFiles("template.txt")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	content["undefined1"] = "オバマ"
	t.Execute(os.Stdout, content)
}

func test2() {
	t, _ := template.ParseFiles("template.txt")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	content["undefined1"] = "オバマ"
	var buffer bytes.Buffer
	t.Execute(&buffer, content)
	fmt.Println(string(buffer.Bytes()))
}

func main() {
	test1();
	test2();
}