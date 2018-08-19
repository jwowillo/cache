package cache

// Getter gets a Value from a Key.
type Getter interface {
	Get(Key) Value
}

// GetterFunc adapts a function to a Getter.
type GetterFunc func(Key) Value

// Get calls the adapted function.
func (f GetterFunc) Get(k Key) Value {
	return f(k)
}

// FallbackGetter is a Getter which gets Values at Keys by checking in a Cache
// first then using the fallback Getter if necessary.
//
// The value returned from the fallback Getter is stored back in the Cache.
type FallbackGetter struct {
	cache  Cache
	getter Getter
}

// NewFallbackGetter which stores in the Cache and falls back to the Getter.
func NewFallbackGetter(c Cache, g Getter) *FallbackGetter {
	return &FallbackGetter{cache: c, getter: g}
}

// Get the Value at the Key by checking the Cache first then using the
// fallback Getter if necessary.
//
// The Value returned from the fallback Getter is stored back in the Cache.
func (g FallbackGetter) Get(k Key) Value {
	v, ok := g.cache.Get(k)
	if !ok {
		v = g.getter.Get(k)
		g.cache.Put(k, v)
	}
	return v
}
