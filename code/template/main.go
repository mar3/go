package main

import "os"
import "fmt"
import "bytes"
import "text/template"

func test1() {
	t, _ := template.ParseFiles("template.txt")
	fields := make(map[string]string)
	fields["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	fields["you"] = "オバマ"
	fields["undefined1"] = "オバマ"
	t.Execute(os.Stdout, fields)
}

func test2() {
	t, _ := template.ParseFiles("template.txt")
	fields := make(map[string]string)
	fields["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	fields["you"] = "オバマ"
	fields["undefined1"] = "オバマ"
	var buffer bytes.Buffer
	t.Execute(&buffer, fields)
	fmt.Println(string(buffer.Bytes()))
}

func main() {
	test1();
	test2();
}