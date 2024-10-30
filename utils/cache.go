package utils

import (
	"crypto/sha1"
	"log"
	"sync"
)

type Shard struct {
  sync.RWMutex
  data map[string]any
}

type ShardMap []*Shard

func NewShardMap(n int) ShardMap {
  shards := make(ShardMap, n)
  
  for i := 0; i < n; i++ {
    shards[i] = &Shard{
      data: make(map[string]any),
    } 
  }

  return shards
}

func (s ShardMap) getShardIndex(key string) int {
  checksum := sha1.Sum([]byte(key))
  hash := int(checksum[0])

  return hash % len(s)
}

func (s ShardMap) getShard(key string) *Shard {
  i := s.getShardIndex(key)
  return s[i]
}

func (s ShardMap) Get(key string) (any, bool) {
  shard := s.getShard(key) 

  shard.RLock()
  defer shard.RUnlock()

  d := shard.data[key]
  return d, d != nil
}

func (s ShardMap) Set(key string, value any) {
  shard := s.getShard(key)

  shard.Lock()
  defer shard.Unlock()

  shard.data[key] = value
}

func (s ShardMap) Delete(key string) {
  shard := s.getShard(key)

  shard.Lock()
  defer shard.Unlock()
  delete(shard.data, key)
}

func (s ShardMap) Contains(key string) bool {
  shard := s.getShard(key)

  shard.RLock()
  defer shard.RUnlock()

  val := shard.data[key]
  return val != nil
}

func (s ShardMap) Keys() []string {
  keys := make([]string, 0)

  mutex := sync.Mutex{}
  wg := sync.WaitGroup{}

  wg.Add(len(s))

  for _, shard := range s {
    go func(s *Shard) {
      s.RLock()

      for k := range s.data {
        mutex.Lock()
        keys = append(keys, k)
        mutex.Unlock()
      }

      s.RUnlock()
      wg.Done()
    }(shard)
  }
  wg.Wait()
  return keys
}

func RunCacheExample() {
  cache := NewShardMap(3)

  cache.Set("a", 1)
  cache.Set("b", 2)
  cache.Set("c", 3)
  
  keys := cache.Keys()
  for k := range keys {
    log.Printf("The key is: %v", k)
  }
}
