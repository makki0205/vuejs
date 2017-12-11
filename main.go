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
	//ウェルカムページ
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	// crosの許可
	api.GET("/makki", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":  "katsuramaki taiki",
			"sex":   "man",
			"email": "llxo2_5oxll@icloud.com",
		})
	})
	r.Run(":3000")
}
