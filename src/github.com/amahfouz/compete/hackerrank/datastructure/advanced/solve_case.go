package advanced

import "fmt"

func solve(a, w []int64) int64 {
	debug := true
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
