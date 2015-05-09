package math

import "testing"

func TestEmptyVector(t *testing.T) {
    v := NewIntVector([]int64 {})
    
    if v.Size() != 0 {
        t.Error("TestEmpty failed. Expected zero length.");
    }
    if v.Items == nil {
        t.Error("TestEmpty failed. Expected non-nill vector.");
    }
}

func TestGet(t *testing.T) {
    v := NewIntVector([]int64 {1, 2, 3})
    
    if v.Size() != 3 {
        t.Error("TestGet failed. Expected length of 3.");
    }
    if v.Items[2] != 3 {
        t.Error("TestGet failed.");
    }
}

func TestMagnitude(t *testing.T) {
	a := NewIntVector([]int64 {3, 0, 4, 0})
	
	if a.Magnitude() != 5.0 {
	    t.Error("TestMagnitude failed. Expected magnitude of 5.");
	}
	    
//	b := NewIntVector([]int64 {3, 0})
}

func TestDot(t *testing.T) {
	a := NewIntVector([]int64 {3, 0, 4, 1})
	b := NewIntVector([]int64 {2, 1, 5, 4})
	
	if a.Dot(*b) != 30 {
		t.Error("TestDot failed.");	    
	}
	if a.Dot(*b) != b.Dot(*a) {
		t.Error("TestDot failed. Expected to be commutitive.");	    
	}	
}