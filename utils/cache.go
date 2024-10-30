package utils

import (
	"log"
	"sync"
)

type Cache struct {
  sync.RWMutex
  data map[string]any
}

func NewCache() Cache {
  return Cache{
    data: make(map[string]any),
  }
}

func (c *Cache) Get(key string) (any, bool) {
  c.RLock()
  defer c.RUnlock()

  d := c.data[key]
  return d, d != nil
}

func (c *Cache) Set(key string, value any) {
  c.Lock()
  defer c.Unlock()

  c.data[key] = value
}

func (c *Cache) Delete(key string) {
  c.Lock()
  defer c.Unlock()
  delete(c.data, key)
}

func (c *Cache) Contains(key string) bool {
  c.RLock()
  defer c.RUnlock()

  val := c.data[key]
  return val != nil
}

func (c *Cache) Keys() []string {
  c.RLock()
  defer c.RUnlock()

  keys := make([]string, 0, len(c.data))
  for k := range c.data {
    keys = append(keys, k)
  }

  return keys
}

func RunCacheExample() {
  cache := NewCache()

  cache.Set("a", 1)
  cache.Set("b", 2)
  cache.Set("c", 3)
  
  keys := cache.Keys()
  for k := range keys {
    log.Printf("The key is: %v", k)
  }
}
