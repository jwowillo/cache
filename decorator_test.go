package cache_test

import (
	"testing"

	"github.com/jwowillo/cache"
)

// TestCompose tests that Compose composes DecoratorFactorys and does so in the
// right order.
func TestCompose(t *testing.T) {
	var order []int
	x := func(cache.Cache) cache.Decorator {
		order = append(order, 1)
		return nil
	}
	y := func(c cache.Cache) cache.Decorator {
		order = append(order, 2)
		return nil
	}
	z := func(cache.Cache) cache.Decorator {
		order = append(order, 3)
		return nil
	}

	cache.Compose()(nil)
	if len(order) != 0 {
		t.Error("order == %v, want %v", order, nil)
	}
	order = nil

	cache.Compose(x)(nil)
	if len(order) != 1 || order[0] != 1 {
		t.Errorf("order = %v, want %v", order, []int{1})
	}
	order = nil

	cache.Compose(x, y)(nil)
	if len(order) != 2 || order[0] != 1 || order[1] != 2 {
		t.Errorf("order = %v, want %v", order, []int{1, 2})
	}
	order = nil

	cache.Compose(x, y, z)(nil)
	if len(order) != 3 || order[0] != 1 || order[1] != 2 || order[2] != 3 {
		t.Errorf("order = %v, want %v", order, []int{1, 2, 3})
	}
	order = nil
}
