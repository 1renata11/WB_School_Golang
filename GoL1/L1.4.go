package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func L14() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, wg)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	<-sigCh
	cancel()
	wg.Wait()
}
