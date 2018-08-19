package decorator

import (
	"time"

	"gopkg.in/jwowillo/cache.v2"
)

// key for last clear.
const key = "last"

// TimeSource returns a time.Time.
type TimeSource func() time.Time

// TimeDecorator is a cache.Decorator which clears all entries after a
// time.Duration passes since the last clear.
type TimeDecorator struct {
	cache cache.Cache

	timeSource TimeSource
	duration   time.Duration

	lastClear cache.Cache
}

// NewTimeDecorator that decorates the cache.Cache and clears all entries after
// a time.Duration passes after the time.Time of the last clear from the
// TimeSource stored in the last-clear cache.Cache.
func NewTimeDecorator(c cache.Cache,
	lastClear cache.Cache, ts TimeSource, d time.Duration) *TimeDecorator {
	lastClear.Put(key, ts())
	return &TimeDecorator{
		cache: c,

		timeSource: ts,
		duration:   d,

		lastClear: lastClear,
	}
}

// NewTimeDecoratorFactory that creates a TimeDecorator with the cache.Cache,
// TimeSource, and time.Duration.
func NewTimeDecoratorFactory(
	lastClear cache.Cache,
	ts TimeSource, d time.Duration) cache.DecoratorFactory {
	return func(c cache.Cache) cache.Decorator {
		return NewTimeDecorator(c, lastClear, ts, d)
	}
}

// Get checks if the current time.Time from the TimeSource is after the
// time.Time the last clear was performed and returns the cache.Value at the
// cache.Key if it wasn't.
//
// Clears the cache.Cache if it was and returns false otherwise.
func (d *TimeDecorator) Get(k cache.Key) (cache.Value, bool) {
	now := d.timeSource()
	lastClear, ok := d.lastClear.Get(key)
	if !ok {
		// This should never happen since a cache.Value is added to the
		// cache.Cache in the constructor and the cache.Cache is never
		// cleared.
		return nil, false
	}
	if now.After(lastClear.(time.Time).Add(d.duration)) {
		d.cache.Clear()
		d.lastClear.Put(key, now)
		return nil, false
	}
	return d.cache.Get(k)
}

// Put calls the decorated cache.Cache's Put.
func (d *TimeDecorator) Put(k cache.Key, v cache.Value) {
	d.cache.Put(k, v)
}

// Delete calls the decorated cache.Cache's Delete.
func (d *TimeDecorator) Delete(k cache.Key) {
	d.cache.Delete(k)
}

// Clear calls the decorated cache.Cache's Clear.
func (d *TimeDecorator) Clear() {
	d.cache.Clear()
}
