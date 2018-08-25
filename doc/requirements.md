# v2.1.0 Requirements

1. Use Go modules for versioning

# v2.0.0 Requirements

1. Move `MemoryCache` into a sub-package called `memory`.
2. Move `LogDecorator`, `ModifiedDecorator`, `ThreadSafeDecorator`, and
   `TimeDecorator` into a sub-package called `decorator`.
3. Move `DefaultLockers`, `DefaultWriter`, `DefaultHasBeenModified`,
   `DefaultTimeSource`, `DefaultModifiedCache`, and `DefaultTimeCache` into a
   sub-package called `standard`.
4. Rename `Fallback` to `Getter` and make an interface, provide a `GetterFunc`
   wrapper, and make `Get` into a struct that implements `Getter` called
   `FallbackGetter`.

# v1.0.0 Requirements

1. Provide an in-memory cache which is thread-safe, logs, and clears itself
   when a key is fetched a provided duration after the last clear
2. Provide an in-memory cache which is thread-safe, logs, and deletes entries
   which have been modified since having been added
