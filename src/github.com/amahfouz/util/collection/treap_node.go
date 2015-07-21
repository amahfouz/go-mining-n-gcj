package collection

type Node struct {
	key, data, priority int64
	left, right, parent *Node
}

type NodeStack struct {
	s    []*Node
	size int
}

// Node methods

func (n *Node) Level() int {
	level := 0
	for n.parent != nil {
		n = n.parent
		level++
	}
	return level
}

func NodeEquals(first, second *Node) bool {
	if (first == nil) != (second == nil) {
		return false
	}

	if first == nil {
		// both are nil
		return true
	}
	// both are non-nil

	return first.data == second.data &&
		first.key == second.key
}

func DeepEquals(first, second *Node) bool {
	if !NodeEquals(first, second) {
		return false
	}

	return first == nil || NodeEquals(first.parent, second.parent) &&
		DeepEquals(first.left, second.left) &&
		DeepEquals(first.right, second.right)
}

// NodeStack methods

func NewNodeStack() NodeStack {
	return NodeStack{make([]*Node, 0, 20), 0}
}

func (stack *NodeStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *NodeStack) Push(n *Node) {
	if stack.size == len(stack.s) {
		stack.s = append(stack.s, n)
	} else {
		stack.s[stack.size] = n
	}

	stack.size++
}

func (stack *NodeStack) Pop() *Node {
	if stack.IsEmpty() {
		panic("Pop called on an empty stack.")
	} else {
		element := stack.s[stack.size-1]
		stack.size--
		return element
	}
}
