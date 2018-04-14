package simplelb

type Request struct {
	fn    func() int
	doneC chan int
}
