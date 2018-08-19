package decorator

import (
	"io"
	"log"

	"gopkg.in/jwowillo/cache.v2"
)

// LogDecorator is a cache.Decorator which logs deleting actions on
// cache.Caches.
type LogDecorator struct {
	cache cache.Cache

	logger *log.Logger
}

// NewLogDecorator that decorates the cache.Cache and writes logs to the
// io.Writer prefixed with name.
func NewLogDecorator(c cache.Cache, w io.Writer, name string) *LogDecorator {
	return &LogDecorator{
		cache: c,

		logger: log.New(w, "cache "+name+": ", log.LstdFlags)}
}

// NewLogDecoratorFactory that creates a LogDecorator with the io.Writer and
// name.
func NewLogDecoratorFactory(w io.Writer, name string) cache.DecoratorFactory {
	return func(c cache.Cache) cache.Decorator {
		return NewLogDecorator(c, w, name)
	}
}

// Get calls the decorated cache.Cache's Get.
func (d LogDecorator) Get(k cache.Key) (cache.Value, bool) {
	return d.cache.Get(k)
}

// Put calls the decorated cache.Cache's Put.
func (d *LogDecorator) Put(k cache.Key, v cache.Value) {
	d.cache.Put(k, v)
}

// Delete logs and calls the decorated cache.Cache's Delete.
func (d *LogDecorator) Delete(k cache.Key) {
	d.logger.Println("delete", k)
	d.cache.Delete(k)
}

// Clear logs and calls the decorated cache.Cache's Clear.
func (d *LogDecorator) Clear() {
	d.logger.Println("clear")
	d.cache.Clear()
}
