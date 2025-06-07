package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]cacheEntry
	Mux sync.Mutex
}

type cacheEntry struct {
	createAt 	time.Time
	val 		[]byte
}
