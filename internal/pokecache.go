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

	var entry_to_Add cacheEntry

	entry_to_add.createdAt = time.Now() 
	entry_to_add.val = data 
  	r.map_result[key_intro] = entry_to_add
}

func (r cache) Get(key_intro string) ([]byte, bool) {
  
	result, exits = r.map_result[key_intro]

	if exists != nil {
		return nil, false
	}
	return result.val, true
}

