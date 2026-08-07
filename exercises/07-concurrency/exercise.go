package concurrency

import (
	"context"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Sum(ctx context.Context, values <-chan int) (int, error) {
	total := 0
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()

		case val, ok := <-values:
			if !ok {
				return total, nil
			}
			total += val
		}
	}
}
