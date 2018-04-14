package simplelb

import "container/heap"

type Balancer struct {
	pool Pool
	done chan *Worker
}

func New(workerCount int) {
	b := Balancer{
		pool: NewPool(workerCount),
		done: make(chan *Worker),
	}

	for i := 0; i < workerCount; i++ {
		w := NewWorker()
		w.work(b.done)
		b.pool = append(b.pool, w)
	}

}

func (b *Balancer) balance(requests chan Request) {
	for {
		select {
		case req := <-requests:
			b.dispatch(req)
		case w := <-b.done:
			b.complete(w)
		}
	}
}

func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker
	w := heap.Pop(&b.pool).(*Worker)
	// send task to it
	w.requests <- req

	w.pending++

	heap.Push(&b.pool, w)

}

func (b *Balancer) complete(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}