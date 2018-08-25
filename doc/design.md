# v2.1.0 Design

* Create a module for the version (1).

# v2.0.0 Design

## `cache` (4)

![`cache` Design](cache_uml.png)

## `cache/memory` (1)

![`cache/memory` Design](memory_uml.png)

## `cache/decorator` (2)

![`cache/decorator` Design](decorator_uml.png)

## `cache/standard`

* `DefaultLockers`, `DefaultWriter`, `DefaultHasBeenModified`,
  `DefaultTimeSource`, `DefaultModifiedCache`, and `DefaultTimeCache` will be
  moved here (3).

# v1.0.0 Design

![`cache` Design](cache.v1_uml.png)

* A default `Cache` will be provided that composes a `ThreadSafeDecorator`,
  `LogDecorator`, and `TimeDecorator` and applies them to a `MemoryCache` (1).
* A default `Cache` will be provided that composes a `ThreadSafeDecorator`,
  `LogDecorator`, and `ModifiedDecorator` and applies them to a
  `MemoryCache` (2).
