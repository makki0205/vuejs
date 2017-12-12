package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {

	})
	r.Run(":3000")
}
