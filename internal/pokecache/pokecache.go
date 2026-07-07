package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	syncMutext *sync.Mutex
	cache      map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		syncMutext: &sync.Mutex{},
		cache:      make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.syncMutext.Lock()
	defer cache.syncMutext.Unlock()
	cache.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.syncMutext.Lock()
	defer cache.syncMutext.Unlock()
	val, ok := cache.cache[key]
	return val.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache *Cache) reap(now time.Time, last time.Duration) {
	cache.syncMutext.Lock()
	defer cache.syncMutext.Unlock()
	for k, v := range cache.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(cache.cache, k)
		}
	}
}
