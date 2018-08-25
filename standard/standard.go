package standard

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/jwowillo/cache/v2"
	"github.com/jwowillo/cache/v2/decorator"
	"github.com/jwowillo/cache/v2/memory"
)

// TimeCache returns a thread-safe cache.Cache that logs to STDOUT with prefix
// name and clears itself every time.Duration.
func TimeCache(name string, duration time.Duration) cache.Cache {
	threadSafe := decorator.NewThreadSafeDecoratorFactory(Lockers())

	log := decorator.NewLogDecoratorFactory(Writer(), name)

	timeCache := decorator.NewThreadSafeDecoratorFactory(Lockers())(
		memory.NewCache())
	time := decorator.NewTimeDecoratorFactory(
		timeCache, TimeSource(), duration)

	return cache.Compose(threadSafe, log, time)(memory.NewCache())
}

// ChangedCache returns a thread-safe cache.Cache that logs to STDOUT with
// prefix name and assumes keys refer to files who are cleared whenever the
// files are changed.
func ChangedCache(name string) cache.Cache {
	threadSafe := decorator.NewThreadSafeDecoratorFactory(Lockers())

	log := decorator.NewLogDecoratorFactory(Writer(), name)

	timeCache := decorator.NewThreadSafeDecoratorFactory(Lockers())(
		memory.NewCache())
	changed := decorator.NewChangedDecoratorFactory(
		timeCache, TimeSource(), HasBeenChanged())

	return cache.Compose(threadSafe, log, changed)(memory.NewCache())
}

// Writer returns os.Stdout.
func Writer() io.Writer {
	return os.Stdout
}

// TimeSource returns time.Now.
func TimeSource() decorator.TimeSource {
	return time.Now
}

// HasBeenChanged returns true if the file at the path has been changed
// since its associated value was stored.
//
// Returns true if the file can't be accessed.
func HasBeenChanged() decorator.HasBeenChanged {
	return func(path cache.Key, last time.Time) bool {
		f, err := os.Stat(string(path))
		if err != nil {
			return true
		}
		return f.ModTime().After(last)
	}
}

// Lockers returns a sync.RWMutex's write sync.Locker and read
// sync.Locker.
func Lockers() (sync.Locker, sync.Locker) {
	m := &sync.RWMutex{}
	return m, m.RLocker()
}
