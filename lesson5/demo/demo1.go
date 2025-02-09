package demo

import (
	"fmt"
	"math/rand"
	"time"
)

func random() int {
	const max int = 100
	return rand.Intn(max)
}

func Goroutine() {
	const size int = 10
	results := []int{}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		go func() {
			results = append(results, random())
		}()
	}
	time.Sleep(time.Second)

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(results[i])
	}
}
