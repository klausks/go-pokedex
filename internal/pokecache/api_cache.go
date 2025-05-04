package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{mutex: &sync.Mutex{}, entries: make(map[string]cacheEntry)}
	ticker := time.NewTicker(interval)
	fmt.Println("test")
	go newCache.reapLoop(ticker.C)
	fmt.Println("Created new cache with interval", interval)
	return &newCache
}

func (c *Cache) Add(key string, value []byte) {
	_, exists := c.entries[key]
	if exists {
		fmt.Println("Key", key, "already exists, no need to add it.")
		return
	}
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: value}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(ch <-chan time.Time) {
	previousTick := time.Now()
	for {
		tick := <-ch
		for key, entry := range c.entries {
			if entry.createdAt.Before(previousTick) {
				delete(c.entries, key)
			}
		}
		previousTick = tick
	}
}
