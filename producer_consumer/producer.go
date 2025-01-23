package main

import (
	"fmt"
	"time"
)

type Producer struct {
	in *chan int
}

func (p Producer) Produce() {
	task := 1
	for {
		*p.in <- task
		fmt.Printf("\nMsg No %d  Produced", task)
		time.Sleep(100 * time.Millisecond)
		task++
	}
}
