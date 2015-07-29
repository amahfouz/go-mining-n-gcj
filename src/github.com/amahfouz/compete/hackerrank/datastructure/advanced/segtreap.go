package advanced

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

import myMath "github.com/amahfouz/util/math"

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

func (t *Treap) Insert(n *Node) {
	n.priority = t.nextRandomPriority()

	if n.lo != n.data {
		panic("Node upper limit must be set to its data.")
	}

	if t.root == nil {
		t.root = n
		return
	}

	var parent *Node = nil
	insertPos := t.root

	for insertPos != nil {
		if n.IsBetter(insertPos) {
			parent = insertPos
			insertPos = insertPos.right
		} else {
			parent = insertPos
			insertPos = insertPos.left
		}
	}

	n.parent = parent

	if n.IsBetter(parent) {
		parent.right = n
	} else {
		parent.left = n
	}
	t.updateMaxRecursively(parent)

	t.rebalance(n)
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

	t.updateMaxRecursively(parent)
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

	t.updateMaxRecursively(parent)
}

func (tree *Treap) updateMaxRecursively(n *Node) {
	for n != nil {
		n.updateLoBound()
		n = n.parent
	}
}

func (t *Treap) isValid(n *Node) bool {
	if n == nil {
		return true
	}

	var leftMin int64 = math.MaxInt64
	if n.left != nil {
		leftMin = n.left.lo
		if n.priority >= n.left.priority ||
			!n.IsBetter(n.left) ||
			!t.isValid(n.left) {
			return false
		}
	}

	var rightMin int64 = math.MaxInt64
	if n.right != nil {
		rightMin = n.right.lo
		if n.priority >= n.right.priority ||
			!n.right.IsBetter(n) ||
			!t.isValid(n.right) {
			return false
		}
	}

	minDescendent := myMath.MinInt64(leftMin, rightMin)
	hasCorrectMin := (n.lo == myMath.MinInt64(n.data, minDescendent))
	return hasCorrectMin
}

func (tree *Treap) findLargestPath(upperBound int64) *Node {
	if tree.root == nil {
		return nil
	}
	return findLargestPath(upperBound, tree.root)
}

func findLargestPath(upperBound int64, node *Node) *Node {
	if node == nil {
		panic("Nil node!")
	}

	if node.lo >= upperBound {
		return nil
	}

	var maxNode *Node = nil

	// go "right" first where nodes with larger value are

	if node.right != nil {
		maxNode = findLargestPath(upperBound, node.right)
	}

	if maxNode != nil {
		return maxNode
	}

	fmt.Printf("Finding : %v, %v", upperBound, node)
	if node.data < upperBound {
		return node
	}

	if node.left != nil {
		return findLargestPath(upperBound, node.left)
	}

	return nil
}
