package queue

import "testing"

func TestQueue(t *testing.T)  {
	var queue = new(Queue)
	var max = 100
	for i:=1;i<=max;i++ {
		queue.EnQueue(i)
	}
	for i:=1;i<=max;i++  {
		item := queue.DeQueue()
		if item != i {
			t.Error("TestQueue failed...")
		}
	}
}