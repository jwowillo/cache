package cache_test

import (
	"testing"

	"github.com/jwowillo/cache"
)

// TestMemoryCacheGetAndPut tests that MemoryCache's Get and Put store and
// retrieve values correctly.
//
// Get and Put are tested together because they depend on each other.
func TestMemoryCacheGetAndPut(t *testing.T) {
	c := cache.NewMemoryCache()

	if _, ok := c.Get("a"); ok {
		t.Errorf("c.Get(%s) = true, want false", "a")
	}

	c.Put("a", 1)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Errorf("c.Get(%s) = %v, false, want 1, true", "a", v)
	}

	c.Put("a", 2)
	if v, ok := c.Get("a"); !ok || v != 2 {
		t.Errorf("c.Get(%s) = %v, false, want 2, true", "a", v)
	}

	c.Put("b", 1)
	if v, ok := c.Get("b"); !ok || v != 1 {
		t.Errorf("c.Get(%s) = %v, false, want 1, true", "b", v)
	}
}

// TestMemoryCacheDelete tests that MemoryCache's Delete deletes elements at the
// correct keys if they exist and does nothing otherwise.
//
// Assumes Get and Put work.
func TestMemoryCacheDelete(t *testing.T) {
	c := cache.NewMemoryCache()

	c.Put("a", 1)
	c.Put("b", 2)

	c.Delete("a")
	if _, ok := c.Get("a"); ok {
		t.Errorf("c.Get(%s) = true, want false", "a")
	}
	if v, ok := c.Get("b"); !ok || v != 2 {
		t.Errorf("c.Get(%s) = %v, false, want 2, true", "b", v)
	}

	// Make sure deleting twice is OK. There's nothing to assert. It's just
	// worthwhile to make sure nothing panics or throws an exception.
	c.Delete("a")

	c.Delete("b")
	if _, ok := c.Get("a"); ok {
		t.Errorf("c.Get(%s) = true, want false", "a")
	}
	if _, ok := c.Get("b"); ok {
		t.Errorf("c.Get(%s) = true, want false", "b")
	}
}

// TestMemoryCacheClear tests that MemoryCache's Clear clears the Cache and does
// nothing if the Cache is empty.
//
// Assumes Get and Put work.
func TestMemoryCacheClear(t *testing.T) {
	c := cache.NewMemoryCache()

	c.Put("a", 1)
	c.Put("b", 2)

	c.Clear()
	if _, ok := c.Get("a"); ok {
		t.Errorf("c.Get(%s) = true, want false", "a")
	}
	if _, ok := c.Get("b"); ok {
		t.Errorf("c.Get(%s) = true, want false", "b")
	}

	// Make sure clearing twice is OK. There's nothing to assert. It's just
	// worthwhile to make sure nothing panics or throws an exception.
	c.Clear()
}
