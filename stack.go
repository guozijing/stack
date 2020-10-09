package stack

import (
	"sync"
)


type ItemStack struct {
	items []int64
	lock  sync.RWMutex
}

func New() *ItemStack {
	s := &ItemStack{}
	return s
}

func (s *ItemStack) Push(t int64) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *ItemStack) Pop() int64 {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1 ]
	s.lock.Unlock()
	return item
}

func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *ItemStack) Size() int {
	return len(s.items)
}

func (s *ItemStack) Top() int64 {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.lock.Unlock()
	return item
}
