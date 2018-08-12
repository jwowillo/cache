package cache

import "sync"

// ThreadSafeDecorator is a Decorator which makes all Cache operations
// thread-safe.
//
// The Decorator assumes that Get doesn't call any Cache functions that modify
// the Cache. You can decorate this with Decorators who's Get modifies the
// Cache. They just shouldn't be decorated by this Decorator.
type ThreadSafeDecorator struct {
	cache Cache

	locker  sync.Locker
	rlocker sync.Locker
}

// NewThreadSafeDecorator that decorates the Cache and calls the write
// sync.Locker for modifying Cache operations and calls the read sync.Locker for
// non-modifying Cache operations.
func NewThreadSafeDecorator(c Cache, l, rl sync.Locker) *ThreadSafeDecorator {
	return &ThreadSafeDecorator{cache: c, locker: l, rlocker: rl}
}

// NewThreadSafeDecoratorFactory that creats a ThreadSafeDecorator with the
// write and read sync.Lockers.
func NewThreadSafeDecoratorFactory(l, rl sync.Locker) DecoratorFactory {
	return func(c Cache) Decorator {
		return NewThreadSafeDecorator(c, l, rl)
	}
}

// Get locks the read sync.Locker and then calls the decorated Cache's Get.
func (d ThreadSafeDecorator) Get(k Key) (Value, bool) {
	d.rlocker.Lock()
	defer d.rlocker.Unlock()
	return d.cache.Get(k)
}

// Put locks the write sync.Locker and then calls the decorated Cache's Put.
func (d *ThreadSafeDecorator) Put(k Key, v Value) {
	d.locker.Lock()
	defer d.locker.Unlock()
	d.cache.Put(k, v)
}

// Delete locks the write sync.Locker and then calls the decorated Cache's
// Delete.
func (d *ThreadSafeDecorator) Delete(k Key) {
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
