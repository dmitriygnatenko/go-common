
## Usage example

```
err := logger.Init(logger.NewConfig(
    logger.WithStdoutLogEnabled(true),
    logger.WithStdoutLogLevel("WARN"),
    logger.WithFileLogEnabled(true),
    logger.WithFileLogLevel("ERROR"),
    logger.WithFilepath("./errors.log"),
))

if err != nil {
    // TODO
}

ctx := context.Background()

ctx = logger.With(ctx, "key1", "value1")

logger.ErrorKV(ctx, "error message", "key2", "value2")
// {"time":"2024-04-19T20:19:17.274684+00:00","level":"ERROR","msg":"error message","key2":"value2","key1":"value1"}
```