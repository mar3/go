package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		response := gin.H{}
		response["message"] = "pong"
		c.JSON(200, response)
	})

	r.Run() // listen and server on 0.0.0.0:8080
}
