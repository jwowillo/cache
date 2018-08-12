package cache

// Fallback gets the Value at a Key.
type Fallback func(Key) Value

// Get the Value at the Key by checking in the Cache first then using the
// Fallback if necessary.
//
// The Value returned from the Fallback is stored back in the Cache.
func Get(c Cache, k Key, fb Fallback) Value {
	v, ok := c.Get(k)
	if !ok {
		v = fb(k)
		c.Put(k, v)
	}
	return v
}
