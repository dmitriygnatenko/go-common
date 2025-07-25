package closer

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var closer = &Closer{}

type Closer struct {
	initOnce  sync.Once
	waitOnce  sync.Once
	mu        sync.Mutex
	functions []func(ctx context.Context) error
	timeout   *time.Duration
	logger    Logger
}

type Logger interface {
	InfoContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

func Init(c Config) {
	closer.initOnce.Do(func() {
		closer.timeout = c.timeout
		closer.logger = c.logger
	})
}

func Add(f func(ctx context.Context) error) {
	closer.mu.Lock()
	defer closer.mu.Unlock()

	closer.functions = append(closer.functions, f)
}

func Wait(ctx context.Context) {
	closer.waitOnce.Do(func() {
		ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		<-ctx.Done()

		// start graceful shutdown
		closeCtx := context.Background()
		var cancel context.CancelFunc
		if closer.timeout != nil {
			closeCtx, cancel = context.WithTimeout(closeCtx, *closer.timeout)
			defer cancel()
		}

		complete := make(chan struct{}, 1)

		go func() {
			for _, f := range closer.functions {
				if err := f(closeCtx); err != nil {
					if closer.logger != nil {
						closer.logger.ErrorContext(closeCtx, "closer: %s", err.Error())
					}
				}
			}

			complete <- struct{}{}
		}()

		select {
		case <-complete:
			if closer.logger != nil {
				closer.logger.InfoContext(closeCtx, "closer: all processes successfully completed")
			}

			break
		case <-closeCtx.Done():
			if closer.logger != nil {
				closer.logger.InfoContext(closeCtx, "closer: completed by timeout")
			}
		}
	})
}
