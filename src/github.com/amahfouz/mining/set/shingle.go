package set

import (
    "github.com/amahfouz/util/collection"
)

func Shingle(s string, length int) collection.StringSet {
    shingles := collection.NewStringSet()
 	for i := 0; i < len(s) - length + 1; i++ {
 	    shingle := s[i : i+length]
 	    shingles.Add(shingle)
 	}   
 	return shingles
}