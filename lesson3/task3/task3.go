package task3

import (
	"sync"
)

type Queue interface {
	Enqueue(element interface{})
	Dequeue() interface{}
}

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(element interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, element)
}

func (q *ConcurrentQueue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}
