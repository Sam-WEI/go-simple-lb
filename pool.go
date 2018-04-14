package simplelb


type Pool []*Worker

func NewPool(count int) Pool {
	return Pool(make([]*Worker, count))
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

func (p Pool) Pop() interface{} {
	p, res := p[:len(p)-1], p[len(p)-1:]
	return res[0]
}
