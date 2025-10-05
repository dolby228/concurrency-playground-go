package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		results <- j * j
		fmt.Printf("worker %d handled %d\n", id, j)
		// time.Sleep(1 * time.Second)
	}
}

func main() {
	const numWorkers = 3
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("result: ", r)
	}
}
