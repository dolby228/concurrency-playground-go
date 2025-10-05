package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int64
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			atomic.AddInt64(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&n))
}
