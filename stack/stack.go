/**
栈
 */
package stack

import "fmt"

type StackItem struct {
	Item interface{}
	Next *StackItem
}

type Stack struct {
	Sp    *StackItem
	Depth int
}

func New() *Stack {
	return new(Stack)
}

func (s *Stack) Push(item interface{}) {
	s.Sp = &StackItem{Item: item, Next: s.Sp}
	s.Depth++
}

func (s *Stack) Pop() interface{} {
	if s.Depth > 0 {
		item := s.Sp.Item
		s.Sp = s.Sp.Next
		s.Depth--
		return item

	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Depth > 0 {
		return s.Sp.Item
	}
	return nil
}

func (s *Stack) Print() {
	sp := s.Sp
	for sp != nil {
		fmt.Println(sp.Item)
		sp = sp.Next
	}
}
