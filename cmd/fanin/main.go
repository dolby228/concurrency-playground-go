package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)


func generator(ctx context.Context, name string, delay time.Duration) <-chan string {
	out := make(chan string)
	//горутина из анонимной функции
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- fmt.Sprintf("%s: %d", name, i):
				time.Sleep(delay + time.Duration(rand.Intn(50))*time.Millisecond)
			}
		}
	}()

	return out
}

func fanIn(ctx context.Context, inputs ...<-chan string) <-chan string {
	out := make(chan string)
	for _, c := range inputs {
		ch := c
		//горутина из анонимной функции
		go func(){
			for {
				select {
				case <-ctx.Done():
					return 
				case v, ok := <-ch:
					if !ok {
						return 
					}
					select {
					case <-ctx.Done():
						return 
					case out <- v:
					}
				}

			}
		}()
	}
	go func ()  {
		<-ctx.Done()
		close(out)
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	a := generator(ctx, "A", 80*time.Millisecond)
	b := generator(ctx, "B", 120*time.Millisecond)

	out := fanIn(ctx, a, b)

	for v := range out {
		fmt.Println(v)
	}
}