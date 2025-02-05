package task1

func Send(ch chan int, num int) {
	ch <- num
}
