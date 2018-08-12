package cache

// MemoryCache which stores Values at Keys in memory.
type MemoryCache struct {
	data map[Key]Value
}

// NewMemoryCache makes an empty MemoryCache.
func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{}
	c.Clear()
	return c
}

// Get the Value at the Key.
//
// Returns true if the Value exists and false otherwise.
func (c MemoryCache) Get(k Key) (Value, bool) {
	v, ok := c.data[k]
	return v, ok
}

// Put the Value at the Key.
func (c *MemoryCache) Put(k Key, v Value) {
	c.data[k] = v
}

// Delete the Key.
func (c *MemoryCache) Delete(k Key) {
	delete(c.data, k)
}

// Clear the MemoryCache.
func (c *MemoryCache) Clear() {
	c.data = make(map[Key]Value)
}
