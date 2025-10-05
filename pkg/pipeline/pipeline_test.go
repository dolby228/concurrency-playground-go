package pipeline

import "testing"


func TestPipeline(t *testing.T) {
	in := Gen(1, 2, 3)
	out := Square(in)

	got := 0
	for v := range out {
		got += v
	}

	if got != (14) {
		t.Fatalf("sum mismatch: got=%d", got)
	}
}