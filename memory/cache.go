// Package memory contains a cache.Cache implementation which stores in memory.
package memory

import "gopkg.in/jwowillo/cache.v2"

// Cache which stores cache.Values at cache.Keys in memory.
type Cache struct {
	data map[cache.Key]cache.Value
}

// NewCache makes an empty Cache.
func NewCache() *Cache {
	c := &Cache{}
	c.Clear()
	return c
}

// Get the cache.Value at the cache.Key.
//
// Returns true if the cache.Value exists and false otherwise.
func (c Cache) Get(k cache.Key) (cache.Value, bool) {
	v, ok := c.data[k]
	return v, ok
}

// Put the cache.Value at the cache.Key.
func (c *Cache) Put(k cache.Key, v cache.Value) {
	c.data[k] = v
}

// Delete the cache.Key.
func (c *Cache) Delete(k cache.Key) {
	delete(c.data, k)
}

// Clear the Cache.
func (c *Cache) Clear() {
	c.data = make(map[cache.Key]cache.Value)
}
