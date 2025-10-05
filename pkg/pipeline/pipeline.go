package pipeline

func Gen(nums ...int) <-chan int {
	out := make(chan int)
	go func(){
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out 
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)

	go func(){
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}