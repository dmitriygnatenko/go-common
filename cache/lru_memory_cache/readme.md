
## Usage example

```
cache := NewCache[string, int64](3)

cache.Set("test1", 100)
cache.Set("test2", 200)
cache.Set("test3", 300)

val1, found1 := cache.Get("test1") // 100, true
val2, found2 := cache.Get("test2") // 200, true
val3, found3 := cache.Get("test3") // 300, true

cache.Set("test4", 400)

val4, found4 := cache.Get("test4") // 400, false
val5, found5 := cache.Get("test1") // 0, false
```