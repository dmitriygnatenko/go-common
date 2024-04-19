
## Usage example

```
dbConn, err := db.NewDB(
    db.NewConfig(
        db.WithDriver("mysql"),
        db.WithUsername("username"),
        db.WithPassword("password"),
        db.WithDatabase("db"),        
    ),
)

if err != nil {
    // TODO
}

err = dbConn.Ping()
```