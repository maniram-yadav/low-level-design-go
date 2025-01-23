package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func main() {
	const NO_CONSUMERS = 10
	runtime.GOMAXPROCS(NO_CONSUMERS)
	in := make(chan int, 1)
	p := Producer{&in}
	c := Consumer{&in, make(chan int, NO_CONSUMERS)}
	go p.Produce()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go c.Consume(ctx)
	wg := &sync.WaitGroup{}
	for i := 0; i < NO_CONSUMERS; i++ {
		go c.Work(i+1, wg)
	}
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	fmt.Println("Process Interrupted")
	cancelFunc()
	wg.Wait()
	fmt.Println("\nMain Function Completed")
	os.Exit(1)
}
