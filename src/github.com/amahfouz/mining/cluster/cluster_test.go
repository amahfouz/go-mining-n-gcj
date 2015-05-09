package cluster

import "testing"
import "github.com/amahfouz/util/test"
import myMath "github.com/amahfouz/util/math"

func TestEmpty(t *testing.T) {
    c := NewCluster()
    test.Assert(t, c.Size() == 0, "Expected zero points.")
}
 
func TestSinglePoint(t *testing.T) {
    c := NewCluster()
    p := myMath.NewPoint(3, 4)
    
    c.Add(p)
    
    test.Assert(t, c.Size() == 1, "Expected one point.")
    test.Assert(t, c.Centroid().Equals(p), "Wrong centroid.")
}


func TestTwoPoints(t *testing.T) {
    c := NewCluster()
	c.Add(myMath.NewPoint(3, 3))
	c.Add(myMath.NewPoint(1, 1))
	
 	test.Assert(t, c.Size() == 2, "Expected two points.")
    test.Assert(t, c.Centroid().Equals(myMath.NewPoint(2, 2)), "Wrong centroid.")	
}

func TestFewPoints(t *testing.T) {
    c := NewCluster()
    c.Add(myMath.NewPoint(1, 1))
    c.Add(myMath.NewPoint(1, 5))
    c.Add(myMath.NewPoint(5, 1))
	c.Add(myMath.NewPoint(5, 5))
	
 	test.Assert(t, c.Size() == 4, "Expected four points.")
    test.Assert(t, c.Centroid().Equals(myMath.NewPoint(3, 3)), "Wrong centroid.")	
}