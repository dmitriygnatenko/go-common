
## Usage example

```
cache, err := NewCache(
    NewConfig(
        WithCapacity(3),
    ),
)

if err != nil {
    // TODO
}

cache.Set("1", "value 1")
cache.Set("2", "value 2")
cache.Set("3", "value 3")

val2, found2 := cache.Get("2") // "value 2", true
val3, found3 := cache.Get("3") // "value 3", true

cache.Set("4", "value 4")

val1, found1 := cache.Get("1") // nil, false
```