package simplelb

import (
	"container/heap"
	"fmt"
)

type Balancer struct {
	pool Pool
	done chan *Worker
}

func NewBalancer(workerCount int) Balancer {
	b := Balancer{
		pool: NewPool(workerCount),
		done: make(chan *Worker),
	}

	for i := 0; i < workerCount; i++ {
		w := NewWorker(fmt.Sprintf("Worker<%v>", i + 1))
		w.work(b.done)
		b.pool = append(b.pool, w)
	}
	return b
}

func (b *Balancer) Balance(requests chan Request) {
	go func() {
		for {
			select {
			case req := <-requests:
				b.dispatch(req)
			case w := <-b.done:
				b.complete(w)
			}
		}
	}()

}

func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker
	w := heap.Pop(&b.pool).(*Worker)
	// send task to it
	w.requests <- req

	fmt.Printf("%v dispatched to %v \n", req.Name, w.name)

	w.pending++

	heap.Push(&b.pool, w)

	b.pool.printWorkerLoad()

}

func (b *Balancer) complete(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}