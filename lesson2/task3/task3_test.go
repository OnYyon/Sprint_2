package task3

import (
	"fmt"
	"slices"
	"testing"
)

func TestSendCh1(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	Send(ch1, ch2)

	vals := []int{}

	for i := 0; i < 3; i++ {
		val := <-ch1
		vals = append(vals, val)
	}
	fmt.Println(vals)
	slices.Sort(vals)

	expected := []int{0, 1, 2}
	if !slices.Equal(vals, expected) {
		t.Fatalf("ch1 values expected: %v, got: %v", expected, vals)
	}
}

func TestSendCh2(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	Send(ch1, ch2)

	vals := []int{}

	for i := 0; i < 3; i++ {
		val := <-ch2
		vals = append(vals, val)
	}

	slices.Sort(vals)

	expected := []int{0, 1, 2}
	if !slices.Equal(vals, expected) {
		t.Fatalf("ch1 values expected: %v, got: %v", expected, vals)
	}
}
