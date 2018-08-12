package cache

import "time"

// key for last clear.
const key = "last"

// TimeSource returns a time.Time.
type TimeSource func() time.Time

// TimeDecorator is a Decorator which clears all entries after a time.Duration
// passes since the last clear.
type TimeDecorator struct {
	cache Cache

	timeSource TimeSource
	duration   time.Duration

	lastClear Cache
}

// NewTimeDecorator that decorates the Cache and clears all entries after a
// time.Duration passes after the time.Time of the last clear from the
// TimeSource stored in the last-clear Cache.
func NewTimeDecorator(c Cache,
	lastClear Cache, ts TimeSource, d time.Duration) *TimeDecorator {
	lastClear.Put(key, ts())
	return &TimeDecorator{
		cache: c,

		timeSource: ts,
		duration:   d,

		lastClear: lastClear,
	}
}

// NewTimeDecoratorFactory that creates a TimeDecorator with the Cache,
// TimeSource, and time.Duration.
func NewTimeDecoratorFactory(
	lastClear Cache, ts TimeSource, d time.Duration) DecoratorFactory {
	return func(c Cache) Decorator {
		return NewTimeDecorator(c, lastClear, ts, d)
	}
}

// Get checks if the current time.Time from the TimeSource is after the time the
// last clear was performed and returns the Value at the Key if it isn't.
//
// Clears the Cache if it is and returns false otherwise.
func (d *TimeDecorator) Get(k Key) (Value, bool) {
	now := d.timeSource()
	lastClear, ok := d.lastClear.Get(key)
	if !ok {
		// This should never happen since a Value is added to the Cache
		// in the constructor and the Cache is never cleared. If it is,
		// return false to trigger the Fallback if used with Get at
		// least.
		return nil, false
	}
	if now.After(lastClear.(time.Time).Add(d.duration)) {
		d.cache.Clear()
		d.lastClear.Put(key, now)
		return nil, false
	}
	return d.cache.Get(k)
}

// Put calls the decorated Cache's Put.
func (d *TimeDecorator) Put(k Key, v Value) {
	d.cache.Put(k, v)
}

// Delete calls the decorated Cache's Delete.
func (d *TimeDecorator) Delete(k Key) {
	d.cache.Delete(k)
}

// Clear calls the decorated Cache's Clear.
func (d *TimeDecorator) Clear() {
	d.cache.Clear()
}
