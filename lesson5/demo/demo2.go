package demo

import (
	"fmt"
	"sync"
	"time"
)

// функция генерирует случайное число в интервале [0, 100)
func GoroutineCorr() {
	const size int = 10
	mx := &sync.Mutex{}
	results := []int{}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		go func() {
			// вызван Lock, поэтому только одна горутина за раз может получить доступ к слайсу
			mx.Lock()
			defer mx.Unlock()
			results = append(results, random())
		}()
	}
	time.Sleep(time.Second)

	// вызван Lock, потому что здесь тоже обращаемся к results
	mx.Lock()
	defer mx.Unlock()
	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(results[i])
	}
}
