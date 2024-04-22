
## Usage example

```
cache := NewCache()

cache.Set("1", "value 1")

val1, found1 := cache.Get("1") // "value 1", true
val2, found2 := cache.Get("2") // nil, false
```