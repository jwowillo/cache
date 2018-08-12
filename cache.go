// Package cache is a caching library which allows functionality to be decorated
// onto basic caches
package cache

// Key in a Cache.
type Key string

// Value at a Key in a Cache.
type Value interface{}

// Cache associates keys with values.
type Cache interface {
	// Get the Value at the Key.
	//
	// Returns false if the Key isn't in the Cache.
	Get(Key) (Value, bool)
	// Put the Value at the Key.
	Put(Key, Value)
	// Delete the Key.
	Delete(Key)
	// Clear the Cache.
	Clear()
}
