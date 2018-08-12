# v1.0.0 Requirements

1. Provide an in-memory cache which is thread-safe, logs, and clears itself
   when a key is fetched a provided duration after the last clear
2. Provide an in-memory cache which is thread-safe, logs, and deletes entries
   which have been modified since having been added
