package task3

// Problem with cycles under 1.22
func Send(ch1, ch2 chan int) {
	for i := 0; i < 3; i++ {
		i := i
		go func(val int) {
			ch1 <- val
		}(i)
		go func(val int) {
			ch2 <- val
		}(i)
	}
}
