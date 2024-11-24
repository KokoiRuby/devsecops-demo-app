package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// load template
	r.LoadHTMLGlob("templates/*")

	// route
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "DevSecOps Demo App",
		})
	})

	// http://localhost:2233
	err := r.Run(":2233")
	if err != nil {
		return
	}
}
