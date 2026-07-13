package concurrency

type Counter struct {
	value int
}

func (c *Counter) Increment() {}

func (c *Counter) Value() int {
	return 0
}
