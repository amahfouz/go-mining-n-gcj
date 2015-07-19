package collection

import "testing"
import "github.com/amahfouz/util/test"

func TestConstruct(t *testing.T) {
	tree := NewTreap()
	test.AssertNil(t, tree.root)
}

func TestNodeEquals(t *testing.T) {
	var n, m *Node

	test.Assert0(t, NodeEquals(n, m))

	n = new(Node)
	test.Assert0(t, !NodeEquals(n, m))

	m = new(Node)
	test.Assert0(t, NodeEquals(n, m))

	n.data = 3
	test.Assert0(t, !NodeEquals(n, m))

	m.data = 3
	test.Assert0(t, NodeEquals(n, m))

	n.key = 1
	test.Assert0(t, !NodeEquals(n, m))

	m.key = 1
	test.Assert0(t, NodeEquals(n, m))

	n.priority = 1000
	test.Assert0(t, !NodeEquals(n, m))

	m.priority = 1000
	test.Assert0(t, NodeEquals(n, m))
}

func TestDeepEquals(t *testing.T) {
	var n, m *Node
	n = new(Node)
	m = new(Node)

	n.key, m.key = 1000, 1000
	n.data, m.data = 3, 3

	var l1, l2 *Node
	l1 = new(Node)
	l2 = new(Node)

	n.left, m.left = l1, l2

	test.Assert0(t, DeepEquals(n, m))

	var r1, r2 *Node
	r1 = new(Node)
	r2 = new(Node)

	n.right, m.right = r1, r2

	test.Assert0(t, DeepEquals(n, m))
}
