package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// tree node
type Node struct {
	index, prevIndex    int64
	data, totalWeight   int64
	lo                  int64
	priority            int64
	left, right, parent *Node
}

// "min" treap where root node is of lowest priority
type Treap struct {
	root               *Node
	nextRandomPriority func() int64
}

// Node methods

func (n *Node) updateLoBound() {
	if n.left == nil && n.right == nil {
		n.lo = n.data
	} else if n.left == nil {
		n.lo = MinInt64(n.right.lo, n.data)
	} else if n.right == nil {
		n.lo = MinInt64(n.left.lo, n.data)
	} else {
		maxDescendent := MinInt64(n.left.lo, n.right.lo)
		n.lo = MinInt64(n.data, maxDescendent)
	}
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

// Treap methods

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

// main

func main() {

	r := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	s.Scan() // scan num cases

	numCases, _ := strconv.Atoi(s.Text())

	for c := 0; c < numCases; c++ {
		s.Scan() // scan number of nodes for this case
		n, _ := strconv.Atoi(s.Text())

		as := make([]int64, n)
		weights := make([]int64, n)

		for i := 0; i < n; i++ {
			s.Scan()
			value, _ := strconv.Atoi(s.Text())
			as[i] = int64(value)
		}
		for i := 0; i < n; i++ {
			s.Scan()
			value, _ := strconv.Atoi(s.Text())
			weights[i] = int64(value)
		}

		solution := solveCase(as, weights)

		fmt.Println(solution)
	}
}
func solveCase(a, w []int64) int64 {
	debug := false
	if len(a) != len(w) {
		panic("Arrays not equal in length.")
	}

	var tree = NewTreap()
	var largestPath *Node

	for i := range a {
		path := tree.findLargestPath(a[i])
		if debug {
			fmt.Printf("Found  : %v\n", path)
		}
		var newPath *Node = new(Node)
		newPath.index = int64(i)
		newPath.data = a[i]
		newPath.lo = a[i]

		if path == nil {
			newPath.prevIndex = -1
			newPath.totalWeight = w[i]
		} else {
			newPath.prevIndex = path.index
			newPath.totalWeight = w[i] + path.totalWeight
		}

		tree.Insert(newPath)

		if largestPath == nil ||
			newPath.totalWeight > largestPath.totalWeight {
			largestPath = newPath

			if debug {
				fmt.Printf("New max: %v\n", largestPath)
			}
		}
	}
	return largestPath.totalWeight
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

	if node.data < upperBound {
		return node
	}

	if node.left != nil {
		return findLargestPath(upperBound, node.left)
	}

	return nil
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}
