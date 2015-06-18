package main

import "fmt"
import "math"

// find the longest increasing subsequence within a given sequence
// see https://en.wikipedia.org/wiki/Longest_increasing_subsequence

func main() {

	// sequence
	X := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}

	n := len(X)

	M := make([]int, n + 1)
	P := make([]int, n)

	// we always have a sub-sequence of length 1, for any non-empty sequence
	M[0] = 0
	P[0] = -1
	L := 1	
	
	// for every node except the first

	for i := 1; i < n; i++ {

		// find where the current element fits

		lo := 1    // recall M[0] is not used
		hi := L

		for lo <= hi {
			mid := int(math.Ceil(float64((lo + hi) / 2)))
			if X[M[mid]] < X[i] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		// now lo is 1 longer than the longest prefix of X[i]
		// that is, lo points to the first element that is 
		// greater than X[i] among those referenced in M

	    M[lo] = i;
	    P[i] = M[lo - 1]
	    
	    if lo > L {
	        L = lo
		}
	    
	} // end for
	
	// gather the sequence backwards
	
	
	fmt.Println(M)
	fmt.Println(P)
	fmt.Println(L)
	seq := make([]int, L)
	
	k := M[L]
	for i := len(seq) - 1; i >= 0; i-- {
	    seq[i] = X[k]
	    k = P[k]
	}
	
	for i := range(seq) {
	    fmt.Println(seq[i])
	}
}
