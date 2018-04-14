package main

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/sam-wei/simplelb"
)

func main() {
	work := make(chan simplelb.Request, 50)
	balancer := simplelb.NewBalancer(6)
	balancer.Balance(work)

	done := make(chan int)
	quit := make(chan bool)

	go startGeneratingNewRequest(work, done)

	workDone := 0
	for {
		select {
		case d := <-done:
			workDone += d
		case <-quit:
			fmt.Println("The work done in total:", workDone)
			return
		}
	}

}

func startGeneratingNewRequest(work chan simplelb.Request, done chan int) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	reqIdx := 0
	for {
		<- ticker.C
		reqIdx++
		idx := reqIdx

		newReq := simplelb.Request{
			Name: fmt.Sprintf("Request<%v>", idx),
			Fn: func() int {
				// working on something
				randDur := time.Duration(rand.Int63n(int64(20 * time.Second)))
				time.Sleep(randDur)
				return int(randDur.Seconds())
			},
			Done: done,
		}

		work <- newReq
	}
}