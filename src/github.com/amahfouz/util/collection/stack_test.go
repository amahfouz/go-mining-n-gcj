package collection

import "testing"
import "github.com/amahfouz/util/test"

func TestEmptyStack(t *testing.T) {
	stack := NewStack(10)
	test.Assert(t, stack.IsEmpty(), "Stack should be empty.")
}

func TestStackPush(t *testing.T) {
	stack := NewStack(10)
	stack.Push(7)
	test.Assert(t, stack.Size() == 1, "Stack should have one element.")
	stack.Push(8)
	test.Assert(t, stack.Size() == 2, "Stack should have two elements.")
}

func TestStackPopFail(t *testing.T) {
	stack := NewStack(10)
	test.AssertFail(t, func() { stack.Pop() }, "Popping from empty stack should fail.")
}

func TestStackPushPushPopPop(t *testing.T) {
	stack := NewStack(10)
	stack.Push(7)
	stack.Push(8)

	element := stack.Pop()
	test.Assert(t, element == 8, "Element pushed last not retrieved on pop.")
	test.Assert(t, stack.Size() == 1, "Expected one element left after pop.")
}

func TestStackPushPopPushPop(t *testing.T) {
	stack := NewStack(10)
	stack.Push(7)
	element := stack.Pop()
	test.Assert(t, element == 7, "Element pushed last not retrieved on pop.")

	stack.Push(8)
	element = stack.Pop()
	test.Assert(t, element == 8, "Unexpected value.")
}
