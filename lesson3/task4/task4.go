package task4

import "sync"

var (
	Buf []int
	mu  sync.Mutex
)

func Write(num int) {
	mu.Lock()
	Buf = append(Buf, num)
	mu.Unlock()
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	val := Buf[0]
	Buf = Buf[1:]
	return val
}
