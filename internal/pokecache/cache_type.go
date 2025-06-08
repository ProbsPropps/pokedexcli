package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data 		map[string]cacheEntry
	Mux 		sync.Mutex
	Duration 	time.Duration
}

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}
