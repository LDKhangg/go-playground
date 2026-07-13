package concurrency

import "context"

type Counter struct {
	value int
}

func (c *Counter) Increment() {}

func (c *Counter) Value() int {
	return 0
}

func Sum(ctx context.Context, values <-chan int) (int, error) {
	return 0, nil
}
