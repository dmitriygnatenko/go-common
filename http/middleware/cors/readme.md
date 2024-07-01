## Usage example

```
corsMiddleware, err := cors.NewCORSMiddleware(
    cors.NewConfig(
        cors.WithOrigin("test.com"),
        cors.WithMethods("GET,POST"),        
    ),
)

if err != nil {
    // TODO
}

srv = &http.Server{
    Addr:    ":8080",
    Handler: corsMiddleware.Handle(mux),
}

```