package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{}

	for i := 0; i < 10000; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Value:", c.Value())

}