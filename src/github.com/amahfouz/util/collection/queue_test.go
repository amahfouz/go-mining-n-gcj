package collection

import "testing"
import "github.com/amahfouz/util/test"

func TestEmptyQueue(t *testing.T) {
	q := NewQueue(10)
	test.Assert(t, q.IsEmpty(), "Queue should be empty.")
}

func TestQueuePush(t *testing.T) {
	queue := NewQueue(10)
	queue.Add(7)
	test.Assert(t, queue.Size() == 1, "Queue should have one element.")
	queue.Add(8)
	test.Assert(t, queue.Size() == 2, "Queue should have two elements.")
}

func TestQueueRemoveFail(t *testing.T) {
	q := NewQueue(10)
	test.AssertFail(t, func() { q.Remove() }, "Removing from empty queue should fail.")
}

func TestQueueAddRemove(t *testing.T) {
	q := NewQueue(10)
	q.Add(7)
	q.Add(8)

	item := q.Remove()
	test.Assert(t, item == 7, "First In is NOT First Out.")
	test.Assert(t, q.Size() == 1, "Queue should have one element.")

	item = q.Remove()
	test.Assert(t, item == 8, "First In is NOT First Out.")
	test.Assert(t, q.Size() == 0, "Queue should have zero elements.")
}
