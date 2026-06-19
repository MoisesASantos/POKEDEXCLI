package pokecache

import (
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

type cache struct {
	map_result map[string]cacheEntry
}

func (r cache) Add(key_intro string, data []byte) {
  map_to_add = r.map_result[key_intro]
  map_to_add.createdAt = time.Now()
  map_to_add.val = data
}

func (r cache) Get(key_intro string) cacheEntry {
  result = r.map_result[key_intro]
}

