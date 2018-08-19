package decorator_test

import (
	"testing"

	"gopkg.in/jwowillo/cache.v2"
)

type MockCache struct {
	GetCalledWith       []cache.Key
	PutKeysCalledWith   []cache.Key
	PutValuesCalledWith []cache.Value
	DeleteCalledWith    []cache.Key
	ClearCalls          int
}

func (c *MockCache) Get(k cache.Key) (cache.Value, bool) {
	c.GetCalledWith = append(c.GetCalledWith, k)
	return nil, false
}

func (c *MockCache) Put(k cache.Key, v cache.Value) {
	c.PutKeysCalledWith = append(c.PutKeysCalledWith, k)
	c.PutValuesCalledWith = append(c.PutValuesCalledWith, v)
}

func (c *MockCache) Delete(k cache.Key) {
	c.DeleteCalledWith = append(c.DeleteCalledWith, k)
}

func (c *MockCache) Clear() {
	c.ClearCalls++
}

// DecoratorTest is a test that makes sure cache.Decorators created by
// cache.DecoratorFactorys decorate cache.Caches correctly.
//
// This only works with cache.DecoratorFactorys that guarantee that each method
// in the cache.Decorator only calls the respective decorated method and only
// calls each method once.
func DecoratorTest(t *testing.T, df cache.DecoratorFactory) {
	DecoratorGetTest(t, df)
	DecoratorPutTest(t, df)
	DecoratorDeleteTest(t, df)
	DecoratorClearTest(t, df)
}

// DecoratorGetTest is a test that makes sure cache.Decorators created by
// cache.DecoratorFactorys decorate cache.Cache's Get correctly.
//
// This only works with cache.DecoratorFactorys that guarantee that the
// cache.Decorator's Get only calls the decorated cache.Cache's Get once and
// calls no other methods.
func DecoratorGetTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Get("k")
	if len(mc.GetCalledWith) != 1 || mc.GetCalledWith[0] != "k" {
		t.Errorf("mc.GetCalledWith = %v, want %v",
			mc.GetCalledWith, []cache.Key{"k"})
	}
}

// DecoratorPutTest is a test that makes sure cache.Decorators created by
// cache.DecoratorFactorys decorate cache.Cache's Put correctly.
//
// This only works with cache.DecoratorFactorys that guarantee that the
// cache.Decorator's Put only calls the decorated cache.Cache's Put once and
// calls no other methods.
func DecoratorPutTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Put("k", 1)
	if len(mc.PutKeysCalledWith) != 1 || mc.PutKeysCalledWith[0] != "k" {
		t.Errorf("mc.PutKeysCalledWith = %v, want %v",
			mc.PutKeysCalledWith, []cache.Key{"k"})
	}
	if len(mc.PutValuesCalledWith) != 1 || mc.PutValuesCalledWith[0] != 1 {
		t.Errorf("mc.PutValuesCalledWith = %v, want %v",
			mc.PutValuesCalledWith, []cache.Value{1})
	}
}

// DecoratorDeleteTest is a test that makes sure cache.Decorators created by
// cache.DecoratorFactorys decorate cache.Cache's Delete correctly.
//
// This only works with cache.DecoratorFactorys that guarantee that the
// cache.Decorator's Get only calls the decorated cache.Cache's Delete once and
// calls no other methods.
func DecoratorDeleteTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Delete("k")
	if len(mc.DeleteCalledWith) != 1 || mc.DeleteCalledWith[0] != "k" {
		t.Errorf("mc.DeleteCalledWith = %v, want %v",
			mc.DeleteCalledWith, []cache.Key{"k"})
	}
}

// DecoratorClearTest is a test that makes sure cache.Decorators created by
// cache.DecoratorFactorys decorate cache.Cache's Clear correctly.
//
// This only works with cache.DecoratorFactorys that guarantee that the
// cache.Decorator's Get only calls the decorated cache.Cache's Clear once and
// calls no other methods.
func DecoratorClearTest(t *testing.T, df cache.DecoratorFactory) {
	mc := MockCache{}
	c := df(&mc)

	c.Clear()
	if mc.ClearCalls != 1 {
		t.Errorf("mc.ClearCalls = %d, want %d",
			mc.ClearCalls, 0)
	}
}
