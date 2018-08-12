package cache

import "time"

// HasBeenModified returns true if the Key has been modified since the
// time.Time.
type HasBeenModified func(Key, time.Time) bool

// ModifiedDecorator is a Decorator which deletes entries that have been
// modified since a time.Time.
type ModifiedDecorator struct {
	cache Cache

	timeSource      TimeSource
	hasBeenModified HasBeenModified

	added Cache
}

// NewModifiedDecorator that decorates the Cache and deletes entries that
// HasBeenModified says has been modified since the time returned from the
// TimeSource.
//
// Times are stored in the added Cache.
func NewModifiedDecorator(c Cache,
	added Cache, ts TimeSource, hbm HasBeenModified) *ModifiedDecorator {
	return &ModifiedDecorator{
		cache: c,

		timeSource:      ts,
		hasBeenModified: hbm,

		added: added,
	}
}

// NewModifiedDecoratorFactory that creates a ModifiedDecorator with the
// Cache, TimeSource, and HasBeenModified.
func NewModifiedDecoratorFactory(
	added Cache,
	ts TimeSource,
	hbm HasBeenModified) DecoratorFactory {
	return func(c Cache) Decorator {
		return NewModifiedDecorator(c, added, ts, hbm)
	}
}

// Get finds the time.Time the Key was originally, calls the decorated Cache's
// Get, checks if the Value at the Key has been modified since that time,
// and returns the Value if it hasn't.
//
// Returns false if the Value doesn't exist. Calls the decorated Cache's Delete
// with the modified Key and returns false if the Key exists but has been
// modified since the time.
func (d *ModifiedDecorator) Get(k Key) (Value, bool) {
	t, ok := d.added.Get(k)
	if !ok {
		return nil, false
	}
	if d.hasBeenModified(k, t.(time.Time)) {
		d.added.Delete(k)
		d.cache.Delete(k)
		return nil, false
	}
	return d.cache.Get(k)
}

// Put notes the time.Time the Key and Value were added calls the decorated
// Cache's Put.
func (d *ModifiedDecorator) Put(k Key, v Value) {
	d.added.Put(k, d.timeSource())
	d.cache.Put(k, v)
}

// Delete calls the decorated Cache's Delete.
func (d *ModifiedDecorator) Delete(k Key) {
	d.added.Delete(k)
	d.cache.Delete(k)
}

// Clear calls the decorated Cache's Clear.
func (d *ModifiedDecorator) Clear() {
	d.added.Clear()
	d.cache.Clear()
}
