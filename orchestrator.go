package main

// Orchestrate woo
func Orchestrate() {
	var workers []*Worker
	buffer := 15
	workChan := make(chan Request, buffer)
	resultReviever := make(chan string)
	var workerID int
	for workerID = 0; workerID < 5; workerID++ {
		workers = append(workers, &Worker{workerID, 0, workChan})
		go working(workers[workerID])
	}
	for i := 0; i < 5; i++ {
		go Requester(workChan, resultReviever, i)
	}
	for {
		select {
		case result := <-resultReviever:
			print(result)
		default:
			if len(workers) >= buffer/2 {
				workers = append(workers, &Worker{workerID, 0, workChan})
				go working(workers[workerID])
				workerID++
			}
		}
	}
}
