
## Usage example

```

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	closer.Init(
	    closer.NewConfig(
            closer.WithTimeout(time.Second * 10),
        )
    )

    // run server 
    
	cancel()
	closer.Wait(ctx)
}
```