package models

import (
	"sync"
	"sync/atomic"
	"time"
)

type Store struct {
	Shorts  int32
	Jackets int32
	Mu      sync.Mutex
}

func (s *Store) RestShorts(orderNumber int) int {
	atomic.AddInt32(&s.Shorts, -1)
	time.Sleep(2 * time.Millisecond)
	return orderNumber
}

func (s *Store) RestJackets(orderNumber int) int {
	atomic.AddInt32(&s.Jackets, -1)
	time.Sleep(2 * time.Millisecond)
	return orderNumber
}
