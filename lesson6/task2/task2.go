package task2

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	c := make(chan int, 1)
	go func(n int) {
		res := fibonacci(n)
		c <- res
	}(n)
	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("Timeout")
	}
}
