package simplelb

type Request struct {
	Name string
	Fn   func() int
	Done chan int
}
