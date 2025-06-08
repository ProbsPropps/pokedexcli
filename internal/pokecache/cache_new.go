package pokecache

import "time"

func NewCache(interval time.Duration) *Cache{
	c := new(Cache)
	c.Data = make(map[string]cacheEntry)
	c.Duration = interval
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.Mux.Lock()
	c.Data[key] = entry
	c.Mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mux.Lock()
	val, exists := c.Data[key]
	c.Mux.Unlock()
	return val.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.Mux.Lock()
			//delete anything older than duration
			for key, entry := range c.Data{
				elapsed := time.Since(entry.createdAt)
				if elapsed > c.Duration{
					delete(c.Data, key)
				}
			}
			c.Mux.Unlock()
		}
	}
}
