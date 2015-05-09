package cluster

import "fmt"
import myMath "github.com/amahfouz/util/math"


// cluster of points along with the sum of their coordinates
// centroid is derivable from the coordinates and number of points

type Cluster struct {
    points myMath.PointSet
    sumX, sumY float64
}

// Cluster methods

func (c *Cluster) Add(p myMath.Point) {
    if c.points.Add(p) {
	    c.sumX += p.X
	    c.sumY += p.Y
    }
}

func (c *Cluster) Remove(p myMath.Point) {
    if c.points.Remove(p) {
	    c.sumX -= p.X
	    c.sumY -= p.Y
    }
}

func (c *Cluster) Centroid() myMath.Point {
    num := c.points.NumElements()
    return myMath.NewPoint(c.sumX / float64(num), c.sumY / float64(num));
}

func (c *Cluster) Size() int {
    return c.points.NumElements()
}

func (c Cluster) String() string {
    setAsStr := fmt.Sprintf("%s", c.points)
	return fmt.Sprintf("Centroid: %s; List{%s}", c.Centroid(), setAsStr); 
}

func NewCluster() Cluster {
	return Cluster{myMath.NewPointSet(), 0.0, 0.0}
} 
