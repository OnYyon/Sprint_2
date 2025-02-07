package task2

import (
	"testing"
)

func TestTask2(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() { ch <- i }()
		val := Receive(ch)
		if val != i {
			t.Fatalf("Expected to receive: %v, got: %v", i, val)
		}
	}
}
