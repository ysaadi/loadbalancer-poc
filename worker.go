package main

// Worker Struct
type Worker struct {
	index   int
	pending int
	ch      chan Request
}

func working(worker *Worker) {
	var req Request
	for {
		req = <-worker.ch
		req.ch <- req.fn(worker)
	}

}
