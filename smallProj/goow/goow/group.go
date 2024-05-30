package goow

import "log"

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	engine      *Engine // all groups share an Engine instance
	middlewares []HandlerFunc
}

func newRouterGroup(prefix string, parent *RouterGroup, engine *Engine) *RouterGroup {
	return &RouterGroup{
		prefix: prefix,
		// parent: parent,
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
