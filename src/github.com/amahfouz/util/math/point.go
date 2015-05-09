package math

import (
    "fmt"
	"bytes"
	"math"
)

// types

type Point struct {
    X, Y float64
}

// Ordered set implementation for small number of points

type PointSet struct {
    points []Point 
}

// Iterator over set of points, returns next Point and 
// 'true' if it is a valid point (within range), and
// 'false' if there are no more points to iterate over
type PointSetIterator func() (Point, bool)   

// functions

func NewPoint(x,y float64) Point {
    return Point{x, y}
}

func (p1 Point) Equals(p2 Point) bool {
    return (p1.X == p2.X) && (p1.Y == p2.Y)
}

func (p Point) String() string {
    return fmt.Sprintf("(%.2f,%.2f)", p.X, p.Y)
}

func (set *PointSet) Add(p Point) bool {
   	if ! set.Contains(p) {
    	set.points = append(set.points, p)
    	return true
	}
   	
   	return false
}

func (set *PointSet) Remove(p Point) bool {
    index := -1
	for i , v := range set.points {
	    if v.Equals(p) {
	        index = i
	        break
	    }
    }
	if index == -1 {
	    return false
	}

	set.points = append(set.points[:index], set.points[index+1:]...)
	return true
}

func (set PointSet) Contains(p Point) bool {
  for _, v := range set.points {
    if v.Equals(p) {
      return true
    }
  }
  return false
}

func (set PointSet) String() string {
    var buffer bytes.Buffer
    for i, p := range set.points {
		buffer.WriteString(p.String())
		if i < len(set.points) - 1 {
		    buffer.WriteString(",");
		}        
    }
    return buffer.String()
}

func (set PointSet) Iterator() PointSetIterator {
    return setIterator(set.points)
} 

func (set PointSet) At(i int) Point {
    return set.points[i]
}

func (set PointSet) NumElements() int {
    return len(set.points)
}

func NewPointSet() PointSet {
	return PointSet{(make([]Point, 0, 10))}
}

func NewPointSetWithPoints(points []Point) PointSet {	
    return PointSet{points} 
}

// non-public methods

func setIterator(items []Point) PointSetIterator {
    index := 0
	return func() (Point, bool) {
		if index >= 0 && index < len(items) {
        	point := items[index]
            index++
            return point, true
		}
        return Point{math.Inf(-1), math.Inf(-1)} ,false
    }    
}