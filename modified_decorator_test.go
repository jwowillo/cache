package cache_test

import (
	"testing"
	"time"

	"github.com/jwowillo/cache"
)

// TestModifiedDecoratorGet tests that ModifiedDecorator's Get deletes stale
// entries without calling the decorated Cache's Get and returns valid entries
// with calling the decorated Cache's Get.
//
// Depends on Put and MemoryCache working.
func TestModifiedDecoratorGet(t *testing.T) {
	var keyCalledWith cache.Key
	var timeCalledWith time.Time
	isModified := false
	ts := &MockTimeSource{}
	ts.Step()
	mc := &MockCache{}
	c := cache.NewModifiedDecoratorFactory(
		cache.NewMemoryCache(),
		ts.Time,
		func(k cache.Key, t time.Time) bool {
			keyCalledWith = k
			timeCalledWith = t
			return isModified
		})(mc)

	if _, ok := c.Get("k"); ok {
		t.Errorf("c.Get(%v) = true, want false", "k")
	}

	c.Put("k", 1)

	c.Get("k")
	if keyCalledWith != "k" {
		t.Errorf("keyCalledWith = %v, want %v", keyCalledWith, "k")
	}
	if !timeCalledWith.Equal(time.Time{}.Add(1)) {
		t.Errorf("timeCalledWith = %v, want %v",
			timeCalledWith, time.Time{}.Add(1))
	}
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []cache.Key{"k"})
	}
	if len(mc.DeleteCalledWith) != 0 {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, []cache.Key{})
	}

	isModified = true

	c.Get("k")
	if len(mc.GetCalledWith) != 1 {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []cache.Key{"k"})
	}
	if len(mc.DeleteCalledWith) != 1 || mc.DeleteCalledWith[0] != "k" {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, []cache.Key{"k"})
	}
}

// TestModifiedDecoratorPut tests that ModifiedDecorator decorates the decorated
// Cache's Put properly.
func TestModifiedDecoratorPut(t *testing.T) {
	ts := &MockTimeSource{}
	f := cache.NewModifiedDecoratorFactory(
		&MockCache{},
		ts.Time,
		func(cache.Key, time.Time) bool { return false })
	DecoratorPutTest(t, f)
}

// TestModifiedDecoratorPut tests that ModifiedDecorator decorates the decorated
// Cache's Delete properly.
func TestModifiedDecoratorDelete(t *testing.T) {
	ts := &MockTimeSource{}
	f := cache.NewModifiedDecoratorFactory(
		&MockCache{},
		ts.Time,
		func(cache.Key, time.Time) bool { return false })
	DecoratorDeleteTest(t, f)
}

// TestModifiedDecoratorPut tests that ModifiedDecorator decorates the decorated
// Cache's Clear properly.
func TestModifiedDecoratorClear(t *testing.T) {
	ts := &MockTimeSource{}
	f := cache.NewModifiedDecoratorFactory(
		&MockCache{},
		ts.Time,
		func(cache.Key, time.Time) bool { return false })
	DecoratorClearTest(t, f)
}
