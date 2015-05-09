package set

import (	
    "fmt"
)

func Jaccard(doc1 string, doc2 string, length int) float64 {
    shingles1 := Shingle(doc1, length)
    shingles2 := Shingle(doc2, length)
    
    fmt.Println(shingles1)
    fmt.Println(shingles2)
    
    intersection := shingles1.Intersect(shingles2)
    union := shingles1.Union(shingles2)

	if union.NumElements() == 0 {
	    panic("Cannot compute similarity of empty sets")
	}

	fmt.Println(union.NumElements())
	fmt.Println(intersection.NumElements())
		
    return float64(intersection.NumElements()) / float64(union.NumElements())
}