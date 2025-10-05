package main

import (
	"fmt"

	"github.com/dolby228/concurrency-playground-go/pkg/pipeline"
)

func main() {
	in := pipeline.Gen(1, 2, 3, 4, 5)
	sq := pipeline.Square(in)

	sum := 0
	for v := range sq {
		sum += v
	}
	fmt.Println("Sum:", sum)
}
