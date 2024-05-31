package goow

import (
	"html/template"
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	// Engine 作为最顶层的分组
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
	// template
	htmpls  *template.Template
	funcMap template.FuncMap
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	// 自己就是一个大组
	engine.RouterGroup = &RouterGroup{engine: engine}
	// main group
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Main function
// 每一个 Socket 来的时候做这个
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	// 按组级来分中间件
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	// 中间件
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmpls = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
