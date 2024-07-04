package goow

import (
	"log"
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	engine      *Engine // all groups share an Engine instance
	middlewares []HandlerFunc
}

func newRouterGroup(prefix string, parent *RouterGroup, engine *Engine) *RouterGroup {
	return &RouterGroup{
		prefix: prefix,
		parent: parent,
		engine: engine,
	}
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := newRouterGroup(group.prefix+prefix, group, group.engine)
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middleware...)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 形成静态文件服务的绝对路径
	absolutePath := path.Join(group.prefix, relativePath)
	// 首先会移除请求URL中的 absolutePath 前缀，
	// 然后使用提供的文件系统 fs 来处理文件服务
	// URL: 0.0.0.0:xxxx/home/static/index.html
	// Path: /static/index.html
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(ctx *Context) {
		file := ctx.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(ctx.Res, ctx.Req)
	}
}

// Static serve static files
// relativePath：相对于 RouterGroup 前缀的静态文件路径。
// rootPath：静态文件的根目录路径。
func (group *RouterGroup) Static(relativePath string, rootPath string) {
	handler := group.createStaticHandler(relativePath, http.Dir(rootPath))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}
