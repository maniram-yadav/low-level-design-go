package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Consumer struct {
	in   *chan int
	jobs chan int
}

func (c Consumer) Work(consumerNo int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range c.jobs {
		time.Sleep(time.Second)
		fmt.Printf("\nJob no %d finished Consumer no %d", consumerNo, job)
	}

}

func (c Consumer) Consume(ctx context.Context) {
	fmt.Println("Inside Consume")
	for {
		select {
		case job := <-*c.in:
			c.jobs <- job
		case <-ctx.Done():
			close(c.jobs)
			return
		}
	}
}
