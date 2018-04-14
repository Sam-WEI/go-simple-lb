package simplelb

type Worker struct {
	requests chan Request
	pending  int
	index    int // index in the heap
}

func NewWorker() *Worker {
	return &Worker{
		requests: make(chan Request),
	}
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <- w.requests
		result := req.fn()
		req.doneC <- result
		done <- w
	}

	// Could run the loop body as a goroutine for parallelism.
}
