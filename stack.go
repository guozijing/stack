package stack

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Item generic.Type

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

func New() *ItemStack {
	s := &ItemStack{}
	s.items = []Item{}
	return s
}

func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1 ]
	s.lock.Unlock()
	return &item
}

func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *ItemStack) Size() int {
	return len(s.items)
}
