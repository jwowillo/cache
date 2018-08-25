package decorator

import (
	"sync"

	"github.com/jwowillo/cache/v2"
)

// ThreadSafeDecorator is a cache.Decorator which makes all cache.Cache
// operations thread-safe.
//
// The cache.Decorator assumes that Get doesn't call any cache.Cache functions
// that modify the cache.Cache. This can be decorated this with cache.Decorators
// who's Get modifies the cache.Cache. They just shouldn't be decorated by this
// cache.Decorator.
type ThreadSafeDecorator struct {
	cache cache.Cache

	locker, rlocker sync.Locker
}

// NewThreadSafeDecorator that decorates the cache.Cache and calls the write
// sync.Locker for modifying cache.Cache operations and calls the read
// sync.Locker for non-modifying cache.Cache operations.
func NewThreadSafeDecorator(
	c cache.Cache, l, rl sync.Locker) *ThreadSafeDecorator {
	return &ThreadSafeDecorator{cache: c, locker: l, rlocker: rl}
}

// NewThreadSafeDecoratorFactory that creates a ThreadSafeDecorator with the
// write and read sync.Lockers.
func NewThreadSafeDecoratorFactory(l, rl sync.Locker) cache.DecoratorFactory {
	return func(c cache.Cache) cache.Decorator {
		return NewThreadSafeDecorator(c, l, rl)
	}
}

// Get locks the read sync.Locker and then calls the decorated cache.Cache's
// Get.
func (d ThreadSafeDecorator) Get(k cache.Key) (cache.Value, bool) {
	d.rlocker.Lock()
	defer d.rlocker.Unlock()
	return d.cache.Get(k)
}

// Put locks the write sync.Locker and then calls the decorated cache.Cache's
// Put.
func (d *ThreadSafeDecorator) Put(k cache.Key, v cache.Value) {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.cache.Put(k, v)
}

// Delete locks the write sync.Locker and then calls the decorated cache.Cache's
// Delete.
func (d *ThreadSafeDecorator) Delete(k cache.Key) {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.cache.Delete(k)
}

// Clear locks the write sync.Locker and then calls the decorated Casche's
// Clear.
func (d *ThreadSafeDecorator) Clear() {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.cache.Clear()
}
