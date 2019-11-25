package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack = new(Stack)
	var max = 1000
	for i:=0;i<=max;i++ {
		stack.Push(i)
	}
	for i:=max;i>=0;i-- {
		item := stack.Pop()
		if item != i {
			t.Error("TestStack failed...")
		}
	}
}