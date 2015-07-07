package collection

// trivial implementation of a queue of int

type Queue struct {
	items []int
	size  int
}

func NewQueue(capacity int) Queue {
	return Queue{make([]int, 0, capacity), 0}
}

func (q Queue) Size() int {
	return q.size
}

func (q Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Add(item int) {
	q.items = append(q.items, item)
	q.size++
}

func (q *Queue) Remove() int {
	if q.IsEmpty() {
		panic("Pop called on an empty queue.")
	} else {
		item := q.items[0]
		q.items = q.items[1:len(q.items)]
		q.size--
		return item
	}
}
