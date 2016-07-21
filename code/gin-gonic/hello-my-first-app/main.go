package main

import "github.com/gin-gonic/gin"

// [/] に対するハンドラーを定義します。
func defaultHandler(c *gin.Context) {

	response := "hi"
	c.String(200, response)
}

// [/api] に対するハンドラーを定義します。
func apiHandler(c *gin.Context) {

	response := gin.H{}
	response["token"] = "aishue9246yhwefhbsc76213lmnzhah1"
	response["message"] = "This is JSON api"
	c.JSON(200, response)
}

func main() {

	r := gin.Default()

	// URL のマッピング
	r.GET("/ping", func(c *gin.Context) {
		response := gin.H{}
		response["message"] = "pong"
		c.JSON(200, response)
	})

	r.GET("/api", apiHandler)

	r.GET("/", defaultHandler)

	// サーバーを起動
	r.Run("0.0.0.0:8080") // listen and server on 0.0.0.0:8080
}
