package main

import (
	"github.com/FeiNiaoBF/Practice-Go/smallProj/goow/goow"
	"net/http"
)

func main() {
	r := goow.New()
	r.GET("/", func(ctx *goow.Context) {
		ctx.HTML(http.StatusOK, "<h1>There Goow</h1>")
	})

	r.GET("/hello", func(ctx *goow.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.POST("/login", func(ctx *goow.Context) {
		ctx.JSON(http.StatusOK, goow.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.GET("/String", func(ctx *goow.Context) {
		ctx.String(http.StatusOK, "hello Goow")
	})

	r.Run(":9999")
}
