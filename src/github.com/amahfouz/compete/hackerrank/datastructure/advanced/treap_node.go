package advanced

import myMath "github.com/amahfouz/util/math"

type Node struct {
	index, prevIndex    int64
	data, totalWeight   int64
	lo                  int64
	priority            int64
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

func (n *Node) updateLoBound() {
	if n.left == nil && n.right == nil {
		n.lo = n.data
	} else if n.left == nil {
		n.lo = myMath.MinInt64(n.right.lo, n.data)
	} else if n.right == nil {
		n.lo = myMath.MinInt64(n.left.lo, n.data)
	} else {
		maxDescendent := myMath.MinInt64(n.left.lo, n.right.lo)
		n.lo = myMath.MinInt64(n.data, maxDescendent)
	}
}

// "static" methods

func NodeEquals(first, second *Node) bool {
	if (first == nil) != (second == nil) {
		return false
	}

	if first == nil {
		// both are nil
		return true
	}
	// both are non-nil

	return first.index == second.index &&
		first.prevIndex == second.prevIndex &&
		first.data == second.data &&
		first.totalWeight == second.totalWeight &&
		first.lo == second.lo
}

func DeepEquals(first, second *Node) bool {
	if !NodeEquals(first, second) {
		return false
	}

	return first == nil ||
		(NodeEquals(first.parent, second.parent) &&
			DeepEquals(first.left, second.left) &&
			DeepEquals(first.right, second.right))
}

// a node is "better" if its totalWeight is more
// if weights are equal, the node with a smaller
// value is better
func (first *Node) IsBetter(second *Node) bool {
	if first == nil || second == nil {
		panic("Cannot compare a nil node")
	}

	if first.totalWeight > second.totalWeight {
		return true
	}
	if first.totalWeight < second.totalWeight {
		return false
	}
	// weights are equal

	if first.data < second.data {
		return true
	}

	if first.data > second.data {
		return false
	}

	// if same data
	if first.index < second.index {
		return true
	}
	return false
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
