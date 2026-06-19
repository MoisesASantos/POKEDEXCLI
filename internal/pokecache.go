package pokecache

import (
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

func (r cacheEntry) Add() int {
  r.createdAt = 
}

timeResult := map[string]cacheEntry
