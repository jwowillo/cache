package decorator_test

import (
	"testing"
	"time"

	"gopkg.in/jwowillo/cache.v2"
	"gopkg.in/jwowillo/cache.v2/decorator"
	"gopkg.in/jwowillo/cache.v2/memory"
)

// TestChangedDecoratorGet tests that ChangedDecorator's Get deletes stale
// entries without calling the decorated cache.Cache's Get and returns valid
// entries with calling the decorated cache.Cache's Get.
//
// Depends on Put and memory.Cache working.
func TestChangedDecoratorGet(t *testing.T) {
	var keyCalledWith cache.Key
	var timeCalledWith time.Time
	isChanged := false
	ts := &MockTimeSource{}
	ts.Step()
	mc := &MockCache{}
	c := decorator.NewChangedDecoratorFactory(
		memory.NewCache(), ts.Time,
		func(k cache.Key, t time.Time) bool {
			keyCalledWith = k
			timeCalledWith = t
			return isChanged
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

	isChanged = true

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

// TestChangedDecoratorPut tests that ChangedDecorator decorates the decorated
// cache.Cache's Put properly.
func TestChangedDecoratorPut(t *testing.T) {
	ts := &MockTimeSource{}
	f := decorator.NewChangedDecoratorFactory(
		&MockCache{},
		ts.Time,
		func(cache.Key, time.Time) bool { return false })
	DecoratorPutTest(t, f)
}

// TestChangedDecoratorPut tests that ChangedDecorator decorates the decorated
// cache.Cache's Delete properly.
func TestChangedDecoratorDelete(t *testing.T) {
	ts := &MockTimeSource{}
	f := decorator.NewChangedDecoratorFactory(
		&MockCache{},
		ts.Time,
		func(cache.Key, time.Time) bool { return false })
	DecoratorDeleteTest(t, f)
}

// TestChangedDecoratorPut tests that ChangedDecorator decorates the decorated
// cache.Cache's Clear properly.
func TestChangedDecoratorClear(t *testing.T) {
	ts := &MockTimeSource{}
	f := decorator.NewChangedDecoratorFactory(
		&MockCache{},
		ts.Time,
		func(cache.Key, time.Time) bool { return false })
	DecoratorClearTest(t, f)
}
