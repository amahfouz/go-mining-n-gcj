package math

import "testing"
import "github.com/amahfouz/util/test"

func TestEmpty(t *testing.T) {
	set := NewPointSet()
	if set.NumElements() != 0 {
	    t.Error("TestEmpty failed.");
	}
}	

func TestInitWithPoints(t *testing.T) {
    points := []Point {NewPoint(1,1), 
    	               NewPoint(10, 10), 
    	               NewPoint(1, 2), 
    	               NewPoint(2,3)}
    pointSet := NewPointSetWithPoints(points)
    
	test.Assert(t, pointSet.NumElements() == 4, "Expected four points.")    
}

func TestPrintPoint(t *testing.T) {
    str := Point{1.00,2.00}.String()
    test.Assert(t, str == "(1.00,2.00)", "Wrong string: " + str)
}

func TestAdd(t *testing.T) {
	set := NewPointSet()
		
	added := set.Add(Point{1, 2})
	
	test.Assert(t, added, "Element not added!")
	test.Assert(t, set.NumElements() == 1, "TestAdd failed. Wrong length.")   
	test.Assert(t, set.Contains(Point{1, 2}), "TestAdd failed. Added element not found.")  
}

func TestAddRepeated(t *testing.T) {
	set := NewPointSet()
		
	set.Add(Point{1, 2})
	added := set.Add(Point{1, 2})
	
	test.Assert(t, ! added, "Element added!")
	test.Assert(t, set.NumElements() == 1, "TestAdd failed. Wrong length.")   
}	

func TestRemoveNotExisting(t *testing.T) {
	set := NewPointSet()
		
	set.Add(Point{1, 2})

	removed := set.Remove(Point{3, 4})
	test.Assert(t, ! removed, "Removed.")
}

func TestRemove(t *testing.T) {
	set := NewPointSet()
		
	set.Add(Point{1, 2})
	set.Add(Point{3, 4})
	set.Add(Point{5, 6})
	
	removed := set.Remove(Point{3, 4})
	test.Assert(t, removed, "Not removed!")
	test.Assert(t, set.NumElements() == 2, "TestRemove failed. Wrong length.")
	test.Assert(t, ! set.Contains(Point{3, 4}), "Element still found.")
}

func TestPrintPointSet(t *testing.T) {
	set := NewPointSet()
		
	set.Add(Point{1, 2})
	set.Add(Point{3, 4})

	str := set.String()
    test.Assert(t, str == "(1.00,2.00),(3.00,4.00)", "Wrong string: " + str)
}

func TestIteratorEmpty(t *testing.T) {
 	set := NewPointSet()
 	iter := set.Iterator()
 	
 	_, valid := iter()
 
 	test.Assert(t, ! valid, "Iterator past end of list!");	   
}

func TestIterator(t *testing.T) {
 	set := NewPointSet()
 	p1 := Point{0, 0}
 	p2 := Point{1, 1}
 	set.Add(p1)
 	set.Add(p2)
 	
 	iter := set.Iterator()
 	
 	next, valid := iter()
 	test.Assert(t, valid, "Iterator not at end of list!");
 	test.Assert(t, next.Equals(p1), "Wrong iterator item.");
 	
 	next, valid = iter()
 	test.Assert(t, valid, "Iterator not at end of list!");
 	test.Assert(t, next.Equals(p2), "Wrong iterator item."); 	   

 	_, valid = iter()
 	test.Assert(t, ! valid, "Iterator past end of list!");	   
}