package main

import (
	"fmt"
	"lesson2/task3"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	task3.Send(ch1, ch2)

	vals := []int{}

	for i := 0; i < 3; i++ {
		val := <-ch1
		vals = append(vals, val)
	}
	fmt.Println(vals)
	vals = []int{}

	for i := 0; i < 3; i++ {
		val := <-ch2
		vals = append(vals, val)
	}
	fmt.Println(vals)
}
