/**
队列
 */
package queue

import "fmt"

type QueueItem struct {
	Item interface{}
	Next *QueueItem
}

type Queue struct {
	Front *QueueItem
	Rear *QueueItem
	Depth int
}

func New() *Queue {
	return new(Queue)
}

func (queue *Queue)EnQueue(item interface{})  {
	q := &QueueItem{Item:item}
	if queue.Depth == 0 {
		queue.Front = q
		queue.Rear = q
	} else {
		queue.Rear.Next = q
		queue.Rear = q
	}
	queue.Depth++
}

func (queue *Queue)DeQueue() interface{} {
	if queue.Depth == 0 {
		return nil
	}
	front := queue.Front
	queue.Front = queue.Front.Next
	queue.Depth--
	return front.Item
}

func (queue *Queue)Print() {
	q := queue.Front
	for q != nil {
		fmt.Println(q.Item)
		q = q.Next
	}
}