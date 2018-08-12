package cache

import (
	"io"
	"log"
)

// LogDecorator is a Decorator which logs deleting actions on Caches.
type LogDecorator struct {
	cache Cache

	logger *log.Logger
}

// NewLogDecorator that decorates the Cache and writes logs to the io.Writer
// prefixed with name.
func NewLogDecorator(c Cache, w io.Writer, name string) *LogDecorator {
	return &LogDecorator{
		cache: c,

		logger: log.New(w, "cache "+name+": ", log.LstdFlags)}
}

// NewLogDecoratorFactory that creates a LogDecorator with the io.Writer and
// name.
func NewLogDecoratorFactory(w io.Writer, name string) DecoratorFactory {
	return func(c Cache) Decorator {
		return NewLogDecorator(c, w, name)
	}
}

// Get calls the decorated Cache's Get.
func (d LogDecorator) Get(k Key) (Value, bool) {
	return d.cache.Get(k)
}

// Put calls the decorated Cache's Put.
func (d *LogDecorator) Put(k Key, v Value) {
	d.cache.Put(k, v)
}

// Delete logs and calls the decorated Cache's Delete.
func (d *LogDecorator) Delete(k Key) {
	d.logger.Println("delete", k)
	d.cache.Delete(k)
}

// Clear logs and calls the decorated Cache's Clear.
func (d *LogDecorator) Clear() {
	d.logger.Println("clear")
	d.cache.Clear()
}
