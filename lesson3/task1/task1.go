package task1

import (
	"errors"
	"sync"
)

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{})}
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	val, ok := s.m[key]
	s.mux.Unlock()
	if !ok {
		return errors.New("Error!")
	}
	return val
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}
