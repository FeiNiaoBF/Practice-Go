package singleflight

import "sync"

// call 代表正在进行中，或者已经结束的请求
// 为了实现请求`去重`，我们需要记录每个 key 对应的请求是否已经在处理中
// 如果是，则其他的请求不需要再次发起，而是等待这个请求结束
// 如果不是，则发起请求，并等待请求结束
// 所以需要 WaitGroup 来实现这个功能
type call struct {
	wg  sync.WaitGroup
	val any
	err error
}

type Group struct {
	mu sync.Mutex // protects m
	m  map[string]*call
}

// Do 代表做一个请求，如果这个请求已经在处理中，则等待这个请求结束
// 所有用户都能收到结果，请求是在服务端阻塞的，
// 等待某一个查询返回结果的，其余请求直接复用这个结果了。
func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
	g.mu.Lock()

	// lazies initialization
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	// 这里的 wg.Done() 一定要在 fn() 之后，否则会出现死锁
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
