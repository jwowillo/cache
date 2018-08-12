package cache

// Decorator is a Cache with extra functionality.
type Decorator Cache

// DecoratorFactory creates a Decorator from a Cache.
type DecoratorFactory func(Cache) Decorator

// Compose DecoratorFactorys into a single DecoratorFactory with the
// DecoratorFactorys called in the order they're passed.
//
// Returns a DecoratorFactory that returns the pased Cache if no
// DecoratorFactorys are passed.
func Compose(dfs ...DecoratorFactory) DecoratorFactory {
	if len(dfs) == 0 {
		return func(c Cache) Decorator { return c }
	}
	for i := len(dfs) - 1; i > 0; i-- {
		inner := dfs[i-1]
		outer := dfs[i]
		dfs[i-1] = func(c Cache) Decorator { return outer(inner(c)) }
	}
	return dfs[0]
}
