## Usage example

```
cache := NewCache[string, int64](3 * time.Second)

cache.Set("test1", -100, nil)

exp := time.Minute
cache.Set("test2", 100, &exp)

val1, found1 := cache.Get("test1") // -100, true
val2, found2 := cache.Get("test2") // 100, true
val3, found3 := cache.Get("test3") // 0, false

time.Sleep(4 * time.Second)

val4, found4 := cache.Get("test1") // 0, false
val5, found5 := cache.Get("test2") // 100, true
```