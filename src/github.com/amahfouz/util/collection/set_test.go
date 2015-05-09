package collection

import (
	"testing"
)

func TestEmpty(t *testing.T) {
    set := NewStringSet()
    if set.NumElements() != 0 {
        t.Error("Expected number of elements to be zero")
    }
}

func TestAdd(t *testing.T) {
 	set := NewStringSet()
 	set.Add("element")
    if set.NumElements() != 1 {
        t.Error("Expected number of elements to be one")
    }
    if ! set.Contains("element") {
        t.Error("Expected the element to be in set")
    }
}

func TestUnion(t *testing.T) {
	set1 := NewStringSet()   
	set2 := NewStringSet()
	
	set1.Add("element1")
	set2.Add("element2")	
	
	set3 := set1.Union(set2)
	
	assert(t, set3.NumElements() == 2, "Expected two elements")
}

func TestIntersection(t *testing.T) {
	set1 := NewStringSet()   
	set2 := NewStringSet()
	
	set1.Add("element1")
	set1.Add("common")
	set2.Add("element2")	
	set2.Add("common")
		
	set3 := set1.Intersect(set2)
	
	assert(t, set3.NumElements() == 1, "Expected one elements")
}

func assert(t *testing.T, condition bool, errMsg string) {
    if ! condition {
        t.Error(errMsg)
    }
}