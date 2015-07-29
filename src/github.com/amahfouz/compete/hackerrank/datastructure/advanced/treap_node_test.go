package advanced

import "testing"
import "github.com/amahfouz/util/test"

func TestLevel(t *testing.T) {
	n := new(Node)

	test.Assert0(t, n.Level() == 0)

	m := new(Node)
	n.parent = m

	test.Assert0(t, n.Level() == 1)
}

func TestNodeEquals(t *testing.T) {
	var n, m *Node

	test.Assert0(t, NodeEquals(n, m))

	n = new(Node)
	test.Assert0(t, !NodeEquals(n, m))

	m = new(Node)
	test.Assert0(t, NodeEquals(n, m))

	n.index = 3
	test.Assert0(t, !NodeEquals(n, m))

	m.index = 3
	test.Assert0(t, NodeEquals(n, m))

	n.prevIndex = 1
	test.Assert0(t, !NodeEquals(n, m))

	m.prevIndex = 1
	test.Assert0(t, NodeEquals(n, m))

	n.totalWeight = 1
	test.Assert0(t, !NodeEquals(n, m))

	m.totalWeight = 1
	test.Assert0(t, NodeEquals(n, m))

	n.lo = 10
	test.Assert0(t, !NodeEquals(n, m))

	m.lo = 10
	test.Assert0(t, NodeEquals(n, m))

	n.data = 10
	test.Assert0(t, !NodeEquals(n, m))

	m.data = 10
	test.Assert0(t, NodeEquals(n, m))
}

func TestDeepEquals(t *testing.T) {
	var n, m *Node
	n = new(Node)
	m = new(Node)

	n.index, m.index = 1000, 1000
	n.totalWeight, m.totalWeight = 3, 3

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

func TestUpdateMax(t *testing.T) {
	n := new(Node)
	l := new(Node)
	r := new(Node)

	test.Assert0(t, n.lo == 0)

	n.data = 10
	n.updateLoBound()
	test.Assert0(t, n.lo == 10)

	l.lo = 8
	n.left = l
	n.updateLoBound()
	test.Assert0(t, n.lo == 8)

	r.lo = 6
	n.right = r
	n.updateLoBound()
	test.Assert0(t, n.lo == 6)

	r.lo = 5
	n.left = nil
	n.updateLoBound()
	test.Assert0(t, n.lo == 5)
}
