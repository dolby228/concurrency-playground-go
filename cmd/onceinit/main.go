package main

import (
	"fmt"
	"sync"
	"time"
)


var (
	once sync.Once
	val string
)

func initSlow() {
	time.Sleep(100*time.Millisecond)
	val = "initialized"
}

func get() string {
	once.Do(initSlow)
	return val
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int){
			defer wg.Done()
			fmt.Printf("G%d: %s\n", id, get())
		}(i)
	}

	wg.Wait()
}
