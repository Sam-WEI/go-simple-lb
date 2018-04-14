package simplelb

import (
	"fmt"
	"strconv"
)

type Pool []*Worker

func NewPool(count int) Pool {
	return Pool(make([]*Worker, 0, count))
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pool) Push(x interface{}) {
	*p = append(*p, x.(*Worker))
}

func (p *Pool) Pop() interface{} {
	res := (*p)[len(*p)-1:]
	*p = (*p)[:len(*p)-1]

	return res[0]
}

func (p Pool) printWorkerLoad() {
	str := "["
	for i, w := range p {
		str += strconv.Itoa(w.pending)
		if i != len(p) - 1 {
			str += ", "
		}
	}

	str += "]"

	fmt.Println(str)
}