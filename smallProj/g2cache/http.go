package g2cache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// http 请求处理其他节点的请求
// 1. 通过http请求获取其他节点的数据
// 2. 通过http请求删除其他节点的数据

const defaultBasePath = "/_g2cache/"

type HTTPPool struct {
	// this peer's base URL, e.g. "https://example.net:8000"
	self string
	// this peer's base path for g2cache HTTP requests
	basePath string
}

// NewHTTPPool creates a new HTTPPool instance.
func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

// Log info with server name
func (p *HTTPPool) Logf(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

// ServeHTTP handles all HTTP requests
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}
	p.Logf("%s || %s", r.Method, r.URL.Path)
	// 约定访问路径格式为 /<basepath>/<groupname>/<key>
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]

	// 获取 group 实例
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	// 根据 key 获取缓存值
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}
