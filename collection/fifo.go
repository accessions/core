package collection

import "sync"

type Queue struct {
	look     sync.Mutex
	elements []interface{}
	size     int
	head     int
	tail     int
	count    int
}

func NewQueue(size int) *Queue {
	return &Queue{
		elements: make([]interface{}, size),
		size:     size,
	}
}

func (q *Queue) Empty() bool {
	q.look.Lock()
	defer q.look.Unlock()
	empty := q.count == 0
	return empty
}

func (q *Queue) Put(element interface{}) {
	q.look.Lock()
	defer q.look.Unlock()

	if q.head == q.tail && q.count > 0 {
		nodes := make([]interface{}, len(q.elements)+q.size)
		copy(nodes, q.elements[q.head:])
		copy(nodes[len(q.elements)-q.head:], q.elements[:q.head])
		q.head = 0
		q.tail = len(q.elements)
		q.elements = nodes
	}

	q.elements[q.tail] = element
	q.tail = (q.tail + 1) % len(q.elements)
	q.count++

}

func (q *Queue) Take() (interface{}, bool) {
	q.look.Lock()
	defer q.look.Unlock()

	if q.count == 0 {
		return nil, false
	}

	element := q.elements[q.head]
	q.head = (q.head + 1) % len(q.elements)
	q.count--

	return element, true
}
