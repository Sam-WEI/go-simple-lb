package simplelb

import "fmt"

type Worker struct {
	name string
	requests chan Request
	pending  int
	index    int // index in the heap
}

func NewWorker(name string) *Worker {
	return &Worker{
		name: name,
		requests: make(chan Request, 100),
	}
}

func (w *Worker) work(done chan *Worker) {
	go func() {
		for req := range w.requests{
			fmt.Printf(">>> %v started %v...\n", w.name, req.Name)
			result := req.Fn()
			fmt.Printf("    <<< %v finished %v...\n", w.name, req.Name)
			req.Done <- result
			done <- w
		}
	}()
	// Could run the loop body as a goroutine for parallelism.
}
