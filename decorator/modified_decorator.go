package decorator

import (
	"time"

	"gopkg.in/jwowillo/cache.v2"
)

// HasBeenChanged returns true if the cache.Key has been changed since the
// time.Time.
type HasBeenChanged func(cache.Key, time.Time) bool

// ChangedDecorator is a cache.Decorator which deletes entries that have been
// changed since a time.Time.
type ChangedDecorator struct {
	cache cache.Cache

	timeSource     TimeSource
	hasBeenChanged HasBeenChanged

	added cache.Cache
}

// NewChangedDecorator that decorates the cache.Cache and deletes entries that
// HasBeenChanged says has been changed since the time.Time returned from the
// TimeSource.
//
// time.Times are stored in the added cache.Cache.
func NewChangedDecorator(
	c cache.Cache,
	a cache.Cache, ts TimeSource, hbm HasBeenChanged) *ChangedDecorator {
	return &ChangedDecorator{
		cache: c,

		timeSource:     ts,
		hasBeenChanged: hbm,

		added: a,
	}
}

// NewChangedDecoratorFactory that creates a ChangedDecorator with the
// cache.Cache, TimeSource, and HasBeenChanged.
func NewChangedDecoratorFactory(
	a cache.Cache,
	ts TimeSource, hbm HasBeenChanged) cache.DecoratorFactory {
	return func(c cache.Cache) cache.Decorator {
		return NewChangedDecorator(c, a, ts, hbm)
	}
}

// Get finds the time.Time the cache.Key was originally added, calls the
// decorated cache.Cache's Get, checks if the cache.Value at the cache.Key has
// been changed since that time, and returns the cache.Value if it hasn't.
//
// Returns false if the cache.Value doesn't exist. Calls the decorated
// cache.Cache's Delete with the changed cache.Key and returns false if the
// cache.Key exists but has been changed since the time.
func (d *ChangedDecorator) Get(k cache.Key) (cache.Value, bool) {
	t, ok := d.added.Get(k)
	if !ok {
		return nil, false
	}
	if d.hasBeenChanged(k, t.(time.Time)) {
		d.added.Delete(k)
		d.cache.Delete(k)
		return nil, false
	}
	return d.cache.Get(k)
}

// Put notes the time.Time the cache.Key and cache.Value were added calls the
// decorated cache.Cache's Put.
func (d *ChangedDecorator) Put(k cache.Key, v cache.Value) {
	d.added.Put(k, d.timeSource())
	d.cache.Put(k, v)
}

// Delete calls the decorated cache.Cache's Delete.
func (d *ChangedDecorator) Delete(k cache.Key) {
	d.added.Delete(k)
	d.cache.Delete(k)
}

// Clear calls the decorated cache.Cache's Clear.
func (d *ChangedDecorator) Clear() {
	d.added.Clear()
	d.cache.Clear()
}
