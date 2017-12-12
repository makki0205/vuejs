package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/vuejs/auth"
)

func main() {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	api := r.Group("/api")
	api.POST("/signin", auth.LoginController)

	authApi := api.Group("")
	authApi.Use(auth.AuthMiddleware)
	authApi.GET("/alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	r.Run(":4000")
}
