package main

import (
	"fmt"
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
	r := goow.Default()
	r.GET("/", func(c *goow.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	r.GET("/panic", func(c *goow.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
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
