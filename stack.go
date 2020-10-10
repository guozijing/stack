package stack

import (
	"sync"
)


type ItemStack struct {
	Items []int
	lock  sync.RWMutex
}

func New() *ItemStack {
	s := &ItemStack{}
	return s
}

func (s *ItemStack) Push(t int) {
	s.lock.Lock()
	s.Items = append(s.Items, t)
	s.lock.Unlock()
}

func (s *ItemStack) Pop() int {
	s.lock.Lock()
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1 ]
	s.lock.Unlock()
	return item
}

func (s *ItemStack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *ItemStack) Size() int {
	return len(s.Items)
}

func (s *ItemStack) Top() int {
	s.lock.Lock()
	item := s.Items[len(s.Items)-1]
	s.lock.Unlock()
	return item
}
