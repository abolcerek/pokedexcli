package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct {
	mapper map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		mapper: make(map[string]cacheEntry),
	}
	
	go newCache.reapLoop(interval)
	return &newCache
}

func (C *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			C.mu.Lock()
			for key , entry := range C.mapper {
				if entry.createdAt.Before(time.Now().Add(-interval)) == true {
					delete(C.mapper, key)
				}
			}
			C.mu.Unlock()
		}
	}
}

func (C *Cache) Add(key string, val []byte) {
	C.mu.Lock()
	defer C.mu.Unlock()
	C.mapper[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (C *Cache) Get(key string) ([]byte, bool) {
	C.mu.Lock()
	defer C.mu.Unlock()
	entry, ok := C.mapper[key]
	if ok != true {
		return nil, false
	}
	return entry.val, true
}