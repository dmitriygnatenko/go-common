## Usage example

```
mux := http.NewServeMux()

config := cors.NewConfig(
    cors.WithOrigin("test.com"),
    cors.WithMethods("GET,POST"),        
)

srv = &http.Server{
    Addr:    ":8080",
    Handler: cors.Handle(config, mux),
}

```