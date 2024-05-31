package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/FeiNiaoBF/Practice-Go/smallProj/goow/goow"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := goow.New()
	r.Use(goow.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("public/template/*")
	r.Static("/assets", "public/static")

	stu1 := &student{Name: "Steve", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *goow.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *goow.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", goow.H{
			"title":  "goow",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *goow.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", goow.H{
			"title": "goow",
			"now":   time.Now().Local().Format("2006-01-02 15:04:05"),
		})
	})

	r.Run(":1314")
}

//func loggerv2() goow.HandlerFunc {
//	return func(c *goow.Context) {
//		// Start timer
//		t := time.Now()
//		// if a server error occurred
//		c.Fail(500, "Internal Server Error")
//		// Calculate resolution time
//		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
//	}
//}
