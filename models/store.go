package models

import (
	"sync"
)

type Store struct {
	Shorts  int
	Jackets int
	Mu      sync.Mutex
}

func (s *Store) RestShorts(orderNumber int) int {
	s.Mu.Lock()
	s.Shorts--
	s.Mu.Unlock()
	return orderNumber * 3
}

func (s *Store) RestJackets(orderNumber int) int {
	s.Mu.Lock()
	s.Jackets--
	s.Mu.Unlock()
	return orderNumber * 3
}
