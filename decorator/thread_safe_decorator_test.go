package decorator_test

import (
	"testing"

	"github.com/jwowillo/cache/v2/decorator"
)

type MockLocker struct {
	LockCount   int
	UnlockCount int
}

func (l *MockLocker) Lock() {
	l.LockCount++
}

func (l *MockLocker) Unlock() {
	l.UnlockCount++
}

// TestThreadSafeDecoratorGetRLocksAndRUnlocks tests that the read sync.Locker
// is called by the cache.Decorator's Get.
func TestThreadSafeDecoratorGetRLocksAndRUnlocks(t *testing.T) {
	l := &MockLocker{}
	rl := &MockLocker{}
	c := decorator.NewThreadSafeDecorator(&MockCache{}, l, rl)
	c.Get("")
	if l.LockCount != 0 {
		t.Errorf("l.LockCount = %d, want %d", l.LockCount, 0)
	}
	if l.UnlockCount != 0 {
		t.Errorf("l.UnlockCount = %d, want %d", l.UnlockCount, 0)
	}
	if rl.LockCount != 1 {
		t.Errorf("rl.LockCount = %d, want %d", l.LockCount, 1)
	}
	if rl.UnlockCount != 1 {
		t.Errorf("rl.UnlockCount = %d, want %d", l.UnlockCount, 1)
	}
}

// TestThreadSafeDecoratorPutLocksAndUnlocks tests that the write sync.Locker is
// called by the cache.Decorator's Put.
func TestThreadSafeDecoratorPutLocksAndUnlocks(t *testing.T) {
	l := &MockLocker{}
	rl := &MockLocker{}
	c := decorator.NewThreadSafeDecorator(&MockCache{}, l, rl)
	c.Put("", nil)
	if l.LockCount != 1 {
		t.Errorf("l.LockCount = %d, want %d", l.LockCount, 1)
	}
	if l.UnlockCount != 1 {
		t.Errorf("l.UnlockCount = %d, want %d", l.UnlockCount, 1)
	}
	if rl.LockCount != 0 {
		t.Errorf("rl.LockCount = %d, want %d", l.LockCount, 0)
	}
	if rl.UnlockCount != 0 {
		t.Errorf("rl.UnlockCount = %d, want %d", l.UnlockCount, 0)
	}
}

// TestThreadSafeDecoratorDeleteLocksAndUnlocks tests that the write sync.Locker
// is called by the cache.Decorator's Delete.
func TestThreadSafeDecoratorDeleteLocksAndUnlocks(t *testing.T) {
	l := &MockLocker{}
	rl := &MockLocker{}
	c := decorator.NewThreadSafeDecorator(&MockCache{}, l, rl)
	c.Delete("")
	if l.LockCount != 1 {
		t.Errorf("l.LockCount = %d, want %d", l.LockCount, 1)
	}
	if l.UnlockCount != 1 {
		t.Errorf("l.UnlockCount = %d, want %d", l.UnlockCount, 1)
	}
	if rl.LockCount != 0 {
		t.Errorf("rl.LockCount = %d, want %d", l.LockCount, 0)
	}
	if rl.UnlockCount != 0 {
		t.Errorf("rl.UnlockCount = %d, want %d", l.UnlockCount, 0)
	}
}

// TestThreadSafeDecoratorClearLocksAndUnlocks tests that the write sync.Locker
// is called by the cache.Decorator's Clear.
func TestThreadSafeDecoratorClearLocksAndUnlocks(t *testing.T) {
	l := &MockLocker{}
	rl := &MockLocker{}
	c := decorator.NewThreadSafeDecorator(&MockCache{}, l, rl)
	c.Clear()
	if l.LockCount != 1 {
		t.Errorf("l.LockCount = %d, want %d", l.LockCount, 1)
	}
	if l.UnlockCount != 1 {
		t.Errorf("l.UnlockCount = %d, want %d", l.UnlockCount, 1)
	}
	if rl.LockCount != 0 {
		t.Errorf("rl.LockCount = %d, want %d", l.LockCount, 0)
	}
	if rl.UnlockCount != 0 {
		t.Errorf("rl.UnlockCount = %d, want %d", l.UnlockCount, 0)
	}
}

// TestThreadSafeDecorator tests that ThreadSafeDecorator decorates properly.
func TestThreadSafeDecorator(t *testing.T) {
	f := decorator.NewThreadSafeDecoratorFactory(
		&MockLocker{}, &MockLocker{})
	DecoratorTest(t, f)
}
