package task2

func Receive(ch chan int) int {
	num := <-ch
	return num
}
