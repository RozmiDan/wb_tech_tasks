package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func withChanCancel(doneCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-doneCh:
			fmt.Println("withChan stopped")
			return
		case <-ticker.C:
			fmt.Println("withChan is working")
		}
	}
}

func withCtxCancel(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("withCtx stopped")
			return
		case <-ticker.C:
			fmt.Println("withCtx is working")
		}
	}
}

func withConditionCancel(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	counter := 0
	for range ticker.C {
		counter++
		fmt.Println("withCond is working")
		if counter == 10 {
			fmt.Println("withCond stopped")
			return
		}
	}
}

func withInputChClose(input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for it := range input {
		_ = it
	}
	fmt.Println("withInputCh stopped")
}

func withGoexitCancel(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("byGoexit working once")
	runtime.Goexit()
}

func main() {
	doneCh := make(chan struct{})
	inputCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	wg.Add(5)

	go withChanCancel(doneCh, wg)
	go withCtxCancel(ctx, wg)
	go withConditionCancel(wg)
	go withInputChClose(inputCh, wg)
	go withGoexitCancel(wg)

	<-ticker.C

	doneCh <- struct{}{}
	close(doneCh)
	close(inputCh)
	cancel()

	wg.Wait()

	fmt.Println("Stop programm")
}
