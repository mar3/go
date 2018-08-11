package main

import "net/http"
import "github.com/labstack/echo"
import "text/template"
import "bytes"

func default_handler(c echo.Context) error {

	t, _ := template.ParseFiles("templates/index.html")
	content := make(map[string]string)
	// content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	// content["you"] = "オバマ"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return c.String(http.StatusOK, string(buffer.Bytes()))
	// return c.String(http.StatusOK, "Hello, World!")
}

func main() {

	e := echo.New()

	// routing
	e.GET("/", default_handler)

	// listenning
	e.Logger.Fatal(e.Start(":8081"))
}
