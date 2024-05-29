package goow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// H is
type H map[string]interface{}

// Context 结构体用于封装HTTP请求和响应的相关信息
type Context struct {
	Res http.ResponseWriter
	Req *http.Request

	Path   string
	Method string
	// eg: [:lang]-{go, py}
	// eg: [*] - {/x, /y, /z}
	Params     map[string]string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Res:    w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 从请求的表单数据中获取指定键的值
func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

// Query 从请求的查询参数中获取指定键的值
func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

// SetHeader 设置响应头
func (ctx *Context) SetHeader(key string, value string) {
	ctx.Res.Header().Set(key, value)
}

func (ctx *Context) Status(status int) {
	ctx.StatusCode = status
}

func (ctx *Context) Param(key string) string {
	return ctx.Params[key]
}

func (ctx *Context) SetDate(Date time.Time) {
	ctx.SetHeader("Date", Date.Format(time.RFC1123))
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.SetDate(time.Now().Local())
	_, _ = ctx.Res.Write([]byte(fmt.Sprintf(format+"\r\n", values...)))
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.SetDate(time.Now().Local())
	_, _ = ctx.Res.Write(data)
}

func (ctx *Context) JSON(code int, body any) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	ctx.SetDate(time.Now().Local())
	encoder := json.NewEncoder(ctx.Res)
	if err := encoder.Encode(body); err != nil {
		http.Error(ctx.Res, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.SetDate(time.Now().Local())
	_, _ = ctx.Res.Write([]byte(html))
}
