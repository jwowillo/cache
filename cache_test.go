package cache_test

import "gopkg.in/jwowillo/cache.v2"

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
