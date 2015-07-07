package collection

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type SegTree struct {
	nodes []int   // hierarchy. Holds a complete bin tree with indices into 'data' array
	data  []int64 // data. Values that determine interval boundaries sorted
	lo    []int   // lower bounds of each interval, inclusive. Parallel to nodes
	hi    []int   // upper bounds of each interval, inclusive. Parallel to nodes

}

func NewSegTree(data []int64) *SegTree {
	if !sort.IsSorted(Int64arr(data)) {
		sort.Sort(Int64arr(data))
	}

	numLeaves := len(data)
	numNodes := 2*numLeaves - 1

	nodes := make([]int, numNodes, numNodes)
	lo := make([]int, numNodes, numNodes)
	hi := make([]int, numNodes, numNodes)

	lo[0] = 0
	hi[0] = numLeaves - 1

	tree := SegTree{nodes, data, lo, hi}
	tree.buildTree()

	fmt.Println(tree.lo)
	return &tree
}

func (tree SegTree) String() string {
	if tree.numLeaves() == 0 {
		return "Empty tree."
	}

	var output bytes.Buffer

	tree.Bft(func(node int) {
		indent := strconv.Itoa(tree.nodeLevel(node))
		output.WriteString(fmt.Sprintf("%"+indent+"d\n", tree.data[node]))
	})

	return output.String()
}

func (tree SegTree) Bft(processNode func(int)) {
	stack := NewStack(tree.numLeaves() / 2)
	stack.Push(0)
	for !stack.IsEmpty() {
		node := stack.Pop()
		processNode(node)
		if !tree.isLeafNode(node) {
			stack.Push(tree.leftChild(node))
			stack.Push(tree.rightChild(node))
		}
	}
}

func (tree SegTree) nodeLevel(node int) int {
	return int(math.Floor(math.Log2(float64(node + 1))))
}

//
// private
//

func (tree SegTree) numLeaves() int {
	return len(tree.data)
}

func (tree SegTree) numNodes() int {
	// n leaf nodes and n-1 internal nodes
	return 2*len(tree.data) - 1
}

func (tree SegTree) isLeafNode(node int) bool {
	return tree.leftChild(node) < 0
}

func (tree SegTree) leftChild(parentIndex int) int {
	childIndex := 2*parentIndex + 1
	if childIndex < tree.numNodes() {
		return childIndex
	} else {
		return -1
	}
}

func (tree SegTree) rightChild(parentIndex int) int {
	childIndex := 2*parentIndex + 1
	if childIndex < tree.numNodes() {
		return childIndex
	} else {
		return -1
	}
}

func (tree SegTree) buildTree() {

	q := NewQueue(tree.numNodes() / 4)
	q.Add(0)

	for !q.IsEmpty() {
		nodeIndex := q.Remove()
		leftDataIndex, rightDataIndex := tree.lo[nodeIndex], tree.hi[nodeIndex]

		fmt.Printf("%d, %d, %d\n", nodeIndex, leftDataIndex, rightDataIndex)
		// if reduced to a single node create a leaf
		if leftDataIndex == rightDataIndex {
			tree.nodes[nodeIndex] = leftDataIndex
			continue
		}

		// not leaf, does not correspond to any data point
		tree.nodes[nodeIndex] = -1

		// split into two. ceil to ensure left tree is at least as high as right tree
		midPoint := int(math.Ceil((float64(rightDataIndex+leftDataIndex) + 1) / 2.0))
		fmt.Println("Mid =", midPoint)

		leftChildIndex := nodeIndex*2 + 1
		rightChildIndex := leftChildIndex + 1

		fmt.Println(leftChildIndex, " ", rightChildIndex)

		tree.lo[leftChildIndex], tree.hi[leftChildIndex] = leftDataIndex, midPoint-1
		tree.lo[rightChildIndex], tree.hi[rightChildIndex] = midPoint, rightDataIndex

		q.Add(leftChildIndex)
		q.Add(rightChildIndex)
	}
}
