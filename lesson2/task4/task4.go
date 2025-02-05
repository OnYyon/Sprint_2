package task4

func Process(nums []int) chan int {
	ch1 := make(chan int, 10)
	for _, i := range nums {
		ch1 <- i
	}
	return ch1
}
