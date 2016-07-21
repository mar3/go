package main

import "os"
import "text/template"

func main() {
	t, _ := template.ParseFiles("template.txt")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	t.Execute(os.Stdout, content)
}
