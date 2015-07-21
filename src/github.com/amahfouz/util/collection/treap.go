package collection

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// "min" treap where root node is of lowest priority
type Treap struct {
	root               *Node
	nextRandomPriority func() int64
}

// methods

func NewTreap() *Treap {
	var randomizer = rand.New(rand.NewSource(time.Now().Unix()))
	randPri := func() int64 {
		return randomizer.Int63n(math.MaxInt64)
	}
	treap := new(Treap)
	treap.nextRandomPriority = randPri
	return treap
}

func (t Treap) NewNode(key, data int64) *Node {
	return &(Node{key, data, t.nextRandomPriority(), nil, nil, nil})
}

func (t *Treap) Add(key, data int64) *Node {

	if t.root == nil {
		t.root = t.NewNode(key, data)
		return t.root
	}

	// append(parents, t.root)

	var parent *Node = nil
	insertPos := t.root

	for insertPos != nil {
		if data < insertPos.data {
			parent = insertPos
			insertPos = insertPos.left
		} else if data > insertPos.data {
			parent = insertPos
			insertPos = insertPos.right
		} else {
			insertPos.key = key
			return nil
		}
	}

	var n *Node = t.NewNode(key, data)
	n.parent = parent

	if n.data < parent.data {
		parent.left = n
	} else {
		parent.right = n
	}

	t.rebalance(n)

	return n
}

func (t *Treap) String() string {
	if t.root == nil {
		return "Empty tree"
	}

	var output bytes.Buffer

	t.Bft(func(node *Node) {
		indent := strconv.Itoa(node.Level() * 4)
		output.WriteString(fmt.Sprintf("%"+indent+"d (%d) %v %v\n", node.data, node.priority, node.left != nil, node.right != nil))
	})

	return output.String()
}

func (tree Treap) Bft(processNode func(*Node)) {
	stack := NewNodeStack()
	stack.Push(tree.root)

	for !stack.IsEmpty() {
		node := stack.Pop()
		processNode(node)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

// private

func (t *Treap) rebalance(n *Node) {
	for n.parent != nil && n.priority < n.parent.priority {
		if n == n.parent.left {
			t.rotateRight(n)
		} else {
			t.rotateLeft(n)
		}
	}
}

func (t *Treap) rotateLeft(n *Node) {
	parent := n.parent
	grandParent := parent.parent

	// rotate

	parent.right = n.left
	n.left = parent

	// update parent links

	if parent.right != nil {
		parent.right.parent = parent
	}
	parent.parent = n
	n.parent = grandParent

	// update grandparent

	if grandParent != nil {
		if grandParent.left == parent {
			grandParent.left = n
		} else {
			grandParent.right = n
		}
	} else {
		t.root = n
	}
}

func (t *Treap) rotateRight(n *Node) {
	parent := n.parent
	grandParent := parent.parent

	// rotate

	parent.left = n.right
	n.right = parent

	// update parent links

	if parent.left != nil {
		parent.left.parent = parent
	}
	parent.parent = n
	n.parent = grandParent

	// update grandparent

	if grandParent != nil {
		if grandParent.left == parent {
			grandParent.left = n
		} else {
			grandParent.right = n
		}
	} else {
		t.root = n
	}
}

func (t *Treap) isValid(n *Node) bool {
	if n == nil {
		return true
	}

	if n.left != nil {
		if n.priority >= n.left.priority ||
			n.left.data >= n.data ||
			!t.isValid(n.left) {
			return false
		}
	}

	if n.right != nil {
		if n.priority >= n.right.priority ||
			n.right.data <= n.data ||
			!t.isValid(n.right) {
			return false
		}
	}

	return true
}
