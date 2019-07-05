package list

import (
	"log"
	"sync"
)

type envList struct {
	items []*Env
	mu    sync.Mutex
}

/**
 * envList
 */
func newEnvList() *envList {
	return &envList{items: make([]*Env, 0, 2)}
}

func (l *envList) add(e *Env) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = append(l.items, e) // append item
}

func (l *envList) remove(e *Env) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for n, item := range l.items {
		if item == e {
			l.items[n] = l.items[0]
			l.items = l.items[1:]
			break
		}
	}
}

func (l *envList) len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.items)
}

/**
 * cache
 */

func (c *Cache) Put(x interface{}) {
	c.mu.Lock()
	if len(c.saved) < cap(c.saved) {
		c.saved = append(c.saved, x)
	} else {
		log.Fatalf(c.name, "if full, you may need to increase pool size")
	}
	c.mu.Unlock()
}

func (c *Cache) Get() interface{} {
	c.mu.Lock()
	n := len(c.saved)
	if n == 0 {
		c.mu.Unlock()
		return c.new()
	}
	x := c.saved[n-1]
	c.saved = c.saved[0 : n-1]
	c.mu.Unlock()
	return x
}

func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.saved = c.saved[:0]
}

func (c *Cache) len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.saved)
}

func (c *Cache) new() *Cache {
	return &Cache{}
}
