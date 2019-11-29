package queue

const Capacity = 100

type CircleQueue struct {
	size int
	data [Capacity]interface{}
	front int
	rear int
}

func (q *CircleQueue)EnQueue(item interface{}) bool {
	if q.size > Capacity {
		return false
	}
	q.size++
	q.data[q.rear] = item
	q.rear = (q.rear+1)%(Capacity-1)	//超过99后，下一个rear回到1
	return true
}

func (q *CircleQueue)DeQueue() interface{} {
	if q.size <= 0 {
		return nil
	}
	q.size--
	item := q.data[q.front]
	q.front = (q.front + 1)%(Capacity - 1)
	return item
}

func (q *CircleQueue)IsEmpty() bool {
	return q.size == 0
}

func (q *CircleQueue) QueueSize() int {
	return q.size
}
