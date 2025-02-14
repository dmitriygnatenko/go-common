
## Usage example

```
cache := NewCache[string, uint64]()

cache.Set("1", 100)

val1, found1 := cache.Get("1") // 100, true
val2, found2 := cache.Get("2") // 0, false
```