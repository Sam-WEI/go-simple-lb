package simplelb

import (
	"time"
	"fmt"
)

func requester(work chan Request) {
	doneC := make(chan int)
	for {
		time.Sleep(5 * time.Second)
		newReq := Request{
			fn: func() int {
				// working on something
				time.Sleep(10 * time.Second)
				return 1
			},
			doneC: doneC,
		}

		work <- newReq

		result := <-doneC

		// do further work with result

		fmt.Println(result)
	}
}