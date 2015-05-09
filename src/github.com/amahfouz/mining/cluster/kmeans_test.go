package cluster

import "testing"
import "fmt"
import "github.com/amahfouz/util/test"
import m "github.com/amahfouz/util/math"

func TestSimpleClustering(t *testing.T) {
    points := []m.Point { m.NewPoint(1,1), 
    	                  m.NewPoint(10, 10), 
    	                  m.NewPoint(1, 2), 
    	                  m.NewPoint(2,3)}
    
    pointSet := m.NewPointSetWithPoints(points)
    	
    set := RunKMeans(2, pointSet, 1)
    
    fmt.Println(set.Clusters[0])
    
    test.Assert(t, len(set.Clusters) == 2, "Expected two clusters.")
    test.Assert(t, set.Clusters[0].Size() == 3, fmt.Sprintf("Expected 3 points but found %d.", set.Clusters[0].Size()))
}

func TestChangeOfAssignment(t *testing.T) {
    points := []m.Point { m.NewPoint(1,1),
        				  m.NewPoint(6,3),
        				  m.NewPoint(4,1),
						  m.NewPoint(3,1), 
    	                  m.NewPoint(2,1), 
    	                  m.NewPoint(6,4),
    	                  m.NewPoint(6,5),
    	                  m.NewPoint(6,6) }
    
    pointSet := m.NewPointSetWithPoints(points)
    	
    set := RunKMeans(2, pointSet, 10)	
    
    fmt.Println(set.Clusters[0])
    fmt.Println(set.Clusters[1])
    
    test.Assert(t, set.Clusters[0].Size() == 4, fmt.Sprintf("Expected 4 points but found %d.", set.Clusters[0].Size()))
    test.Assert(t, set.Clusters[1].Size() == 4, fmt.Sprintf("Expected 3 points but found %d.", set.Clusters[1].Size()))


   
}