package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Request struct
type Request struct {
	ch chan int
	fn func(worker *Worker) int
}

// Requester is used to generate requests to the orchestrator
func Requester(work chan<- Request, resultChan chan string, id int) {
	ch := make(chan int)
	for {
		select {
		case <-ch:
			resultChan <- fmt.Sprintf("a request returned from requester %d", id)
		default:
			time.Sleep(time.Duration(rand.Int63n(10000)) * time.Millisecond)
			resultChan <- fmt.Sprintf("creating a request from requester %d", id)
			work <- Request{ch, (*Worker).workFn}
		}
	}
}

func (w *Worker) workFn() int {
	time.Sleep(time.Duration(rand.Int63n(10000)) * time.Millisecond)
	fmt.Printf("working... %d", w.index)
	return 0
}
