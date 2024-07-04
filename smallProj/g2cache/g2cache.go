package g2cache

import (
	"fmt"
	"log"
	"sync"
)

// Getter 用于从数据源获取数据
type Getter interface {
	// Get 回调函数
	Get(key string) ([]byte, error)
}

// GetterFunc 用于实现 Getter 接口
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex              // 读写锁
	groups = make(map[string]*Group) // group cluster
)

// A Group is a cache namespace and associated data loaded spread over
// 一个 Group 就是一个节点
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		log.Println("[G2Cache] hit")
		return v, nil
	}
	// 缓存未命中
	// 从数据源获取数据
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	// 从数据源获取数据
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}
// populateCache 将源数据添加到缓存 mainCache
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}