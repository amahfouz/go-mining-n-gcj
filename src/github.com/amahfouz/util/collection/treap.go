package collection

import (
	"math"
	"math/rand"
	"time"
)

// types

type Node struct {
	key, data, priority int64
	left, right, parent *Node
}

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
	return n
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
		first.key == second.key &&
		first.priority == second.priority
}

func DeepEquals(first, second *Node) bool {
	if !NodeEquals(first, second) {
		return false
	}

	return first == nil || NodeEquals(first.parent, second.parent) &&
		DeepEquals(first.left, second.left) &&
		DeepEquals(first.right, second.right)
}
