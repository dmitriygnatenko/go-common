## Usage example

```
cache := NewCache(
    NewConfig(
        memoryCache.WithCleanupInterval(30*time.Minute),
        memoryCache.WithExpiration(time.Hour),
    ),
)

cache.Set("1", "value 1", nil)

exp := 12*time.Hour
cache.Set("2", "value 2", &exp)
	
val1, found1 := cache.Get("1") // "value 1", true
val3, found3 := cache.Get("3") // nil, false
```