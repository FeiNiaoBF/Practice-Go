package main

import (
	"net/http"

	"github.com/FeiNiaoBF/Practice-Go/smallProj/goow/goow"
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

	r.GET("/hello/:name", func(ctx *goow.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})

	r.GET("/assets/*filepath", func(ctx *goow.Context) {
		ctx.JSON(http.StatusOK, goow.H{"filepath": ctx.Param("filepath")})
	})

	r.GET("/String", func(ctx *goow.Context) {
		ctx.String(http.StatusOK, "hello Goow")
	})

	// GROUP BY
	v1 := r.Group("/g1")
	{
		v1.GET("/", func(c *goow.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *goow.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/g2")
	{
		v2.GET("/hello/:name", func(c *goow.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *goow.Context) {
			c.JSON(http.StatusOK, goow.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}


	r.Run(":9999")
}
