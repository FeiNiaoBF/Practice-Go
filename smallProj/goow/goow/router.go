package goow

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	// pattern root
	urlPattern map[string]*PatRoot
	// map Handler function
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		urlPattern: make(map[string]*PatRoot),
		handlers:   make(map[string]HandlerFunc),
	}
}

func parseToPattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			// Only one * is allowed
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {

	parts := parseToPattern(pattern)
	key := fmt.Sprint(method + "-" + pattern)
	// GET -- tree
	// POST -- tree
	_, ok := r.urlPattern[method]
	if !ok {
		r.urlPattern[method] = NewPatRoot()
	}

	r.urlPattern[method].insert(pattern, parts)
	r.handlers[key] = handler
	//fmt.Printf("urlPattern: %+v\n", r.urlPattern)
	log.Printf("Route %4s - %s", method, pattern)
}

// getRoute 查找匹配的路由规则
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parseToPattern(path)
	params := make(map[string]string)
	root, ok := r.urlPattern[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts)

	if n != nil {
		parts := parseToPattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern
		ctx.handlers = append(ctx.handlers, r.handlers[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	// 开始做handler
	ctx.Next()
}
