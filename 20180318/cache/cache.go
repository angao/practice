package cache

import "sync"

type Cache struct {
	mu sync.Mutex
	values map[interface{}]interface{}
}

func (c Cache) Lookup(i interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.values[i]
}
