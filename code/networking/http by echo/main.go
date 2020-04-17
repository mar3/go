package main

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func default_handler(c echo.Context) error {

	t, _ := template.ParseFiles("templates/index.html")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return c.HTML(http.StatusOK, string(buffer.Bytes()))
}

func hello1_handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func form1_get_handler(c echo.Context) error {
	t, _ := template.ParseFiles("templates/form1.html")
	content := make(map[string]string)
	content["text1"] = "オバマ1"
	content["text2"] = "オバマ2"
	content["text3"] = "オバマ3"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return c.HTML(http.StatusOK, string(buffer.Bytes()))
}

func form1_post_handler(c echo.Context) error {
	t, _ := template.ParseFiles("templates/form1_result.html")
	content := make(map[string]string)
	content["text1"] = "オバマ1"
	content["text2"] = "オバマ2"
	content["text3"] = "オバマ3"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return c.HTML(http.StatusOK, string(buffer.Bytes()))
}

func main() {

	e := echo.New()

	// middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	// routing
	e.GET("/", default_handler)
	e.GET("/hello1", hello1_handler)

	e.GET("/form1", form1_get_handler)
	e.POST("/form1", form1_post_handler)

	// listenning
	e.Logger.Fatal(e.Start(":8081"))
}
